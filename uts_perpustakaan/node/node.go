package node

type Buku struct {
	ID        int
	Judul     string
	Pengarang string
	Tahun     int
	Halaman   int
}

type NodeBukuBerlisensi struct {
	Data Buku
	Next *NodeBukuBerlisensi
}

type NodeBukuTidakBerlisensi struct {
	IDEntry     int
	IDReferensi int
	Next        *NodeBukuTidakBerlisensi
}
