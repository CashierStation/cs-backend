package enum

type KaryawanRole string

const (
	Owner    KaryawanRole = "owner"
	Karyawan KaryawanRole = "karyawan"
)

func (e KaryawanRole) String() string {
	switch e {
	case Owner:
		return "owner"
	case Karyawan:
		return "karyawan"
	default:
		return "unknown"
	}
}
