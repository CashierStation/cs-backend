package enum

type KaryawanRole string

const (
	Owner    KaryawanRole = "owner"
	Karyawan KaryawanRole = "karyawan"
)

func (e KaryawanRole) String() string {
	switch e {
	case Owner:
		return "Owner"
	case Karyawan:
		return "Karyawan"
	default:
		return "Unknown"
	}
}
