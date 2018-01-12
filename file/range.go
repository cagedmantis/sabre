package file

import "fmt"

type RangeSystem struct {
	ranges []*Range
}

func NewRangeSystem() *RangeSystem {
	return &RangeSystem{
		ranges: make([]*Range, 0, 100),
	}
}

func (rs *RangeSystem) GetRange(fileOffset uint64) (*Range, error) {
	rangeIdx := fileOffset % defaultRange
	if len(rs.ranges) > int(rangeIdx) {
		return nil, fmt.Errorf("range does not exist")
	}

	return rs.ranges[rangeIdx], nil
}

//TODO Seems like the range needs to know where the last range ended. Consider where these
// numbers are being calculated
func (rs *RangeSystem) CreateRange(fileOffset uint64, chunkID uint64, chunkOffset uint64) (*Range, error) {
	rangeIdx := fileOffset % defaultRange
	if int(rangeIdx) < len(rs.ranges) {
		return nil, fmt.Errorf("range already exists")
	}

	r := NewRange(chunkID, chunkOffset, defaultRange, fileOffset)

	rs.ranges = append(rs.ranges, r)

	return r, nil
}

type Range struct {
	chunkID     uint64
	chunkOffset uint64
	size        uint64
	fileOffset  uint64
}

func NewRange(chunkID uint64, chunkOffset uint64, size uint64, fileOffset uint64) *Range {
	return &Range{
		chunkID:     chunkID,
		chunkOffset: chunkOffset,
		size:        size,
		fileOffset:  fileOffset,
	}
}
