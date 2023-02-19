package v1

// import (
// 	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/service"
// 	"github.com/mrsubudei/chat-bot-backend/appointment-service/pkg/logger"
// )

// type translationRoutes struct {
// 	s service.Service
// 	l logger.Interface
// }

// // func newTranslationRoutes(handler *gin.RouterGroup, s service.Service, l logger.Interface) {
// // 	r := &translationRoutes{s, l}

// // 	h := handler.Group("/events")
// // 	{
// // 		h.POST("/create", r.CreateEvents)
// // 	}
// // }

// // func (r *translationRoutes) CreateEvents(c *gin.Context) {

// // 	schedule := entity.Schedule{}

// // 	if err := c.ShouldBindJSON(&schedule); err != nil {
// // 		r.l.Error(err, "http - v1 - CreateEvents")
// // 		errorResponse(c, http.StatusBadRequest, "invalid request body")

// // 		return
// // 	}

// // err := r.s.CreateSchedule(c.Request.Context(), schedule)
// // if err != nil {
// // r.l.Error(err, "http - v1 - CreateEvents")
// // errorResponse(c, http.StatusInternalServerError, "events service problems")
// //
// // return
// // }
// //
// // c.Writer.WriteHeader(200)
// // }
