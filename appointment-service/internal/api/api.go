package api

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/entity"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/repository"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/pkg/logger"
	pb "github.com/mrsubudei/chat-bot-backend/appointment-service/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AppointmentServer struct {
	pb.UnimplementedAppointmentServer
	repo repository.Events
	l    logger.Interface
}

func NewAppointmentServer(repo repository.Events, l logger.Interface) *AppointmentServer {
	return &AppointmentServer{
		repo: repo,
		l:    l,
	}
}

func (as *AppointmentServer) CreateDoctor(ctx context.Context,
	in *pb.DoctorSingle) (*pb.Empty, error) {

	doctor := entity.Doctor{
		Name:    in.Value.Name,
		Surname: in.Value.Surname,
		Phone:   in.Value.Phone,
	}

	err := as.repo.StoreDoctor(ctx, doctor)
	if err != nil {
		if strings.Contains(err.Error(), DuplicateErrMsg) {
			return &pb.Empty{}, status.Error(codes.AlreadyExists,
				entity.ErrDoctorAlreadyExists.Error())
		}
		as.l.Error(fmt.Errorf("api - CreateDoctor - StoreDoctor: %w", err))
		return &pb.Empty{}, status.Error(codes.Internal, InternalErr)
	}

	return &pb.Empty{}, nil
}

func (as *AppointmentServer) GetDoctor(ctx context.Context,
	in *pb.IdRequest) (*pb.DoctorSingle, error) {

	found, err := as.repo.GetDoctor(ctx, in.Id)
	if err != nil {
		if errors.Is(err, entity.ErrDoctorDoesNotExist) {
			return &pb.DoctorSingle{}, status.Error(codes.NotFound,
				entity.ErrDoctorDoesNotExist.Error())
		}
		as.l.Error(fmt.Errorf("api - GetDoctor - GetDoctor: %w", err))
		return &pb.DoctorSingle{}, status.Error(codes.Internal,
			InternalErr)
	}

	doctor := &pb.DoctorSingle{
		Value: &pb.Doctor{
			Id:      found.Id,
			Name:    found.Name,
			Surname: found.Surname,
			Phone:   found.Phone,
		},
	}

	return doctor, nil
}

func (as *AppointmentServer) UpdateDoctor(ctx context.Context,
	in *pb.DoctorSingle) (*pb.DoctorSingle, error) {

	toUpdate := entity.Doctor{
		Id:      in.Value.Id,
		Name:    in.Value.Name,
		Surname: in.Value.Surname,
		Phone:   in.Value.Phone,
	}

	updated, err := as.repo.UpdateDoctor(ctx, toUpdate)
	if err != nil {
		if errors.Is(err, entity.ErrDoctorDoesNotExist) {
			return &pb.DoctorSingle{}, status.Error(codes.NotFound,
				entity.ErrDoctorDoesNotExist.Error())
		}
		as.l.Error(fmt.Errorf("api - UpdateDoctor - UpdateDoctor: %w", err))
		return &pb.DoctorSingle{}, status.Error(codes.Internal,
			InternalErr)
	}

	doctor := &pb.DoctorSingle{
		Value: &pb.Doctor{
			Id:      updated.Id,
			Name:    updated.Name,
			Surname: updated.Surname,
			Phone:   updated.Phone,
		},
	}

	return doctor, nil
}

func (as *AppointmentServer) DeleteDoctor(ctx context.Context,
	in *pb.IdRequest) (*pb.Empty, error) {

	err := as.repo.DeleteDoctor(ctx, in.Id)
	if err != nil {
		if strings.Contains(err.Error(), NoRowsAffected) {
			return &pb.Empty{}, status.Error(codes.NotFound,
				entity.ErrDoctorDoesNotExist.Error())
		}
		as.l.Error(fmt.Errorf("api - DeleteDoctor - DeleteDoctor: %w", err))
		return &pb.Empty{}, status.Error(codes.Internal, InternalErr)
	}

	return &pb.Empty{}, nil
}

func (as *AppointmentServer) GetAllDoctors(ctx context.Context,
	in *pb.Empty) (*pb.DoctorMultiple, error) {

	found, err := as.repo.FetchDoctors(ctx)
	if err != nil {
		as.l.Error(fmt.Errorf("api - GetAllDoctors - GetAllDoctors: %w", err))
		return &pb.DoctorMultiple{}, status.Error(codes.Internal, InternalErr)
	}

	doctors := []*pb.Doctor{}

	for i := 0; i < len(found); i++ {
		doctor := &pb.Doctor{
			Id:      found[i].Id,
			Name:    found[i].Name,
			Surname: found[i].Surname,
			Phone:   found[i].Phone,
		}
		doctors = append(doctors, doctor)
	}

	return &pb.DoctorMultiple{
		Value: doctors,
	}, nil
}

func (as *AppointmentServer) CreateSchedule(ctx context.Context,
	in *pb.ScheduleSingle) (*pb.Empty, error) {

	dayEvents := []entity.Event{}
	var err error
	var existEvent time.Time

	if err := as.checkValues(in); err != nil {
		return &pb.Empty{}, err
	}

	firstDay := in.Value.FirstDay.AsTime()
	lastDay := in.Value.LastDay.AsTime().AddDate(0, 0, 1)
	startTime := in.Value.StartTime.AsTime().Format(DateAndTimeFormat)[11:]
	endTime := in.Value.EndTime.AsTime().Format(DateAndTimeFormat)[11:]
	startBreak := in.Value.StartBreak.AsTime().Format(DateAndTimeFormat)[11:]
	endBreak := in.Value.EndBreak.AsTime().Format(DateAndTimeFormat)[11:]
	increase := time.Duration(int(in.Value.EventDurationMinutes) * int(time.Minute))

	for firstDay.Before(lastDay) {
		date := firstDay.Format(DateAndTimeFormat)[:11]
		starts, err := time.Parse(DateAndTimeFormat, date+startTime)
		if err != nil {
			as.l.Error(fmt.Errorf("api - CreateSchedule - Parse #1: %w", err))
			return &pb.Empty{},
				status.Error(codes.Internal, CanNotParseTime)
		}

		ends, err := time.Parse(DateAndTimeFormat, date+endTime)
		if err != nil {
			as.l.Error(fmt.Errorf("api - CreateSchedule - Parse #2: %w", err))
			return &pb.Empty{},
				status.Error(codes.Internal, CanNotParseTime)
		}

		startsBreak, err := time.Parse(DateAndTimeFormat, date+startBreak)
		if err != nil {
			as.l.Error(fmt.Errorf("api - CreateSchedule - Parse #3: %w", err))
			return &pb.Empty{},
				status.Error(codes.Internal, CanNotParseTime)
		}

		endsBreak, err := time.Parse(DateAndTimeFormat, date+endBreak)
		if err != nil {
			as.l.Error(fmt.Errorf("api - CreateSchedule - Parse #4: %w", err))
			return &pb.Empty{},
				status.Error(codes.Internal, CanNotParseTime)
		}

		was := true
		for i := starts; i.Before(ends); i = i.Add(increase) {
			if was {
				if i.Equal(startsBreak) || i.After(startsBreak) {
					i = endsBreak
					was = false
				}
			}
			event := entity.Event{}
			event.StartsAt = i
			event.EndsAt = i.Add(increase)
			for u := 0; u < len(in.Value.DoctorIds); u++ {
				event.DoctorId = in.Value.DoctorIds[u]
				dayEvents = append(dayEvents, event)
			}
		}

		firstDay = firstDay.AddDate(0, 0, 1)
	}

	existEvent, err = as.repo.StoreSchedule(ctx, dayEvents)
	if err != nil {
		if errors.Is(err, entity.ErrEventAlreadyExists) {
			str := existEvent.Format(DateAndTimeFormat)
			return &pb.Empty{},
				status.Error(codes.AlreadyExists,
					entity.ErrEventAlreadyExists.Error()+": "+str)
		}
		as.l.Error(fmt.Errorf("api - CreateSchedule - StoreSchedule: %w", err))
		return &pb.Empty{},
			status.Error(codes.Internal, InternalErr)
	}

	return &pb.Empty{}, nil
}

func (as *AppointmentServer) checkValues(in *pb.ScheduleSingle) error {
	isFirstDayZero := in.Value.FirstDay.AsTime().IsZero()
	isLastDayZero := in.Value.LastDay.AsTime().IsZero()
	st := in.Value.StartTime.AsTime()
	isStartTimeZero := st.Equal(time.Date(st.Year(), st.Month(),
		st.Day(), 0, 0, 0, 0, st.Location()))
	et := in.Value.EndTime.AsTime()
	isEndTimeZero := et.Equal(time.Date(et.Year(), et.Month(),
		et.Day(), 0, 0, 0, 0, et.Location()))
	sb := in.Value.StartBreak.AsTime()
	isStartBreakZero := sb.Equal(time.Date(sb.Year(), sb.Month(),
		sb.Day(), 0, 0, 0, 0, sb.Location()))
	eb := in.Value.EndBreak.AsTime()
	isEndBreakZero := eb.Equal(time.Date(eb.Year(), eb.Month(),
		eb.Day(), 0, 0, 0, 0, eb.Location()))

	switch {
	case isFirstDayZero == true:
		return status.Error(codes.InvalidArgument,
			RequestZeroValue+": FirstDay")
	case isLastDayZero == true:
		return status.Error(codes.InvalidArgument,
			RequestZeroValue+": LastDay")
	case isStartTimeZero == true:
		return status.Error(codes.InvalidArgument,
			RequestZeroValue+": StartTime")
	case isEndTimeZero == true:
		return status.Error(codes.InvalidArgument,
			RequestZeroValue+": EndTime")
	case isStartBreakZero == true:
		return status.Error(codes.InvalidArgument,
			RequestZeroValue+": StartBreak")
	case isEndBreakZero == true:
		return status.Error(codes.InvalidArgument,
			RequestZeroValue+": EndBreak")

	}

	return nil
}

func (as *AppointmentServer) GetOpenEventsByDoctor(ctx context.Context,
	in *pb.IdRequest) (*pb.EventMultiple, error) {

	found, err := as.repo.FetchOpenEventsByDoctor(ctx, in.Id)
	if err != nil {
		as.l.Error(fmt.Errorf("api - GetOpenEventsByDoctor - "+
			"FetchOpenEventsByDoctor: %w", err))
		return &pb.EventMultiple{}, status.Error(codes.Internal, InternalErr)
	}

	events := convertEvents(found)

	return &pb.EventMultiple{
		Value: events,
	}, nil
}

func (as *AppointmentServer) GetReservedEventsByDoctor(ctx context.Context,
	in *pb.IdRequest) (*pb.EventMultiple, error) {

	found, err := as.repo.FetchReservedEventsByDoctor(ctx, in.Id)
	if err != nil {
		as.l.Error(fmt.Errorf("api - GetReservedEventsByDoctor - "+
			"FetchReservedEventsByDoctor: %w", err))
		return &pb.EventMultiple{}, status.Error(codes.Internal, InternalErr)
	}

	events := convertEvents(found)

	return &pb.EventMultiple{
		Value: events,
	}, nil
}

func (as *AppointmentServer) GetReservedEventsByClient(ctx context.Context,
	in *pb.IdRequest) (*pb.EventMultiple, error) {

	found, err := as.repo.FetchReservedEventsByClient(ctx, in.Id)
	if err != nil {
		as.l.Error(fmt.Errorf("api - GetReservedEventsByClient - "+
			"FetchReservedEventsByClient: %w", err))
		return &pb.EventMultiple{}, status.Error(codes.Internal, InternalErr)
	}

	events := convertEvents(found)

	return &pb.EventMultiple{
		Value: events,
	}, nil
}

func (as *AppointmentServer) GetAllEventsByClient(ctx context.Context,
	in *pb.IdRequest) (*pb.EventMultiple, error) {

	found, err := as.repo.FetchAllEventsByClient(ctx, in.Id)
	if err != nil {
		as.l.Error(fmt.Errorf("api - GetAllEventsByClient - "+
			"FetchAllEventsByClient: %w", err))
		return &pb.EventMultiple{}, status.Error(codes.Internal, InternalErr)
	}

	events := convertEvents(found)

	return &pb.EventMultiple{
		Value: events,
	}, nil
}

func (as *AppointmentServer) RegisterToEvent(ctx context.Context,
	in *pb.EventSingle) (*pb.EventSingle, error) {

	toUpdate := entity.Event{
		Id:       in.Value.Id,
		ClientId: in.Value.ClientId,
		DoctorId: in.Value.DoctorId,
		StartsAt: in.Value.StartsAt.AsTime(),
		EndsAt:   in.Value.EndsAt.AsTime(),
	}

	updated, err := as.repo.UpdateEvent(ctx, toUpdate)
	if err != nil {
		if errors.Is(err, entity.ErrEventDoesNotExist) {
			return &pb.EventSingle{}, status.Error(codes.NotFound,
				entity.ErrEventDoesNotExist.Error())
		} else if errors.Is(err, entity.ErrEventAlreadyReserved) {
			return &pb.EventSingle{}, status.Error(codes.Canceled,
				entity.ErrEventAlreadyReserved.Error())
		}
		as.l.Error(fmt.Errorf("api - RegisterToEvent - UpdateEvent: %w", err))
		return &pb.EventSingle{}, status.Error(codes.Internal, InternalErr)
	}

	return &pb.EventSingle{
		Value: &pb.Event{
			Id:       updated.Id,
			ClientId: updated.ClientId,
			DoctorId: updated.DoctorId,
			StartsAt: timestamppb.New(updated.StartsAt),
			EndsAt:   timestamppb.New(updated.EndsAt),
		},
	}, nil
}

func (as *AppointmentServer) UnregisterEvent(ctx context.Context,
	in *pb.EventSingle) (*pb.Empty, error) {

	toUpdate := entity.Event{
		Id:       in.Value.Id,
		ClientId: in.Value.ClientId,
		DoctorId: in.Value.DoctorId,
		StartsAt: in.Value.StartsAt.AsTime(),
		EndsAt:   in.Value.EndsAt.AsTime(),
	}

	err := as.repo.ClearEvent(ctx, toUpdate)
	if err != nil {
		if errors.Is(err, entity.ErrEventDoesNotExist) {
			return &pb.Empty{}, status.Error(codes.NotFound,
				entity.ErrEventDoesNotExist.Error())
		} else if errors.Is(err, entity.ErrEventIsNotReserved) {
			return &pb.Empty{}, status.Error(codes.Canceled,
				entity.ErrEventIsNotReserved.Error())
		}
		as.l.Error(fmt.Errorf("api - UnregisterEvent - ClearEvent: %w", err))
		return &pb.Empty{}, status.Error(codes.Internal, InternalErr)
	}

	return &pb.Empty{}, nil
}

func convertEvents(sl []entity.Event) []*pb.Event {
	events := []*pb.Event{}

	for i := 0; i < len(sl); i++ {
		event := &pb.Event{
			Id:       sl[i].Id,
			ClientId: sl[i].ClientId,
			DoctorId: sl[i].DoctorId,
			StartsAt: timestamppb.New(sl[i].StartsAt),
			EndsAt:   timestamppb.New(sl[i].EndsAt),
		}
		events = append(events, event)
	}

	return events
}
