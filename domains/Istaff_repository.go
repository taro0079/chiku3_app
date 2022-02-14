package domains

type IStaffRepository interface {
	dbConnect()
	Save()
}
