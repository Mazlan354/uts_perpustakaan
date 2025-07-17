package model

import "uts_perpustakaan/node"

var BukuBerlisensiHead *node.NodeBukuBerlisensi
var BukuTidakBerlisensiHead *node.NodeBukuTidakBerlisensi

func TambahBukuBerlisensi(bk node.Buku) {
	nodeBuku := &node.NodeBukuBerlisensi{Data: bk}
	if BukuBerlisensiHead == nil {
		BukuBerlisensiHead = nodeBuku
	} else {
		temp := BukuBerlisensiHead
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = nodeBuku
	}
}

func CariBukuBerlisensi(id int) *node.Buku {
	temp := BukuBerlisensiHead
	for temp != nil {
		if temp.Data.ID == id {
			return &temp.Data
		}
		temp = temp.Next
	}
	return nil
}

func TambahBukuTidakBerlisensi(idEntry int, idReferensi int) bool {
	if CariBukuBerlisensi(idReferensi) == nil {
		return false
	}
	nodeBuku := &node.NodeBukuTidakBerlisensi{
		IDEntry:     idEntry,
		IDReferensi: idReferensi,
	}
	if BukuTidakBerlisensiHead == nil {
		BukuTidakBerlisensiHead = nodeBuku
	} else {
		temp := BukuTidakBerlisensiHead
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = nodeBuku
	}
	return true
}

func GetBukuTidakBerlisensi() *node.NodeBukuTidakBerlisensi {
	return BukuTidakBerlisensiHead
}
