package system

type RouterGroup struct {
	CalendarEventRouter
	CalendarGroupRouter
	CalendarUserRouter
	CalendarBaseRouter
	CalendarDBRouter
	SiteMonitoringRouter
	PictureStorageRouter
}
