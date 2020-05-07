package repositories

import (
	"github.com/jinzhu/gorm"
)

type UpVotesRepo interface {
	GetUpVoteCounts(siteId string, pageId string) (count int, err error)
}

type upVotesRepo struct {
	db *gorm.DB
}

func NewUpVotesRepo(db *gorm.DB) UpVotesRepo {
	return &upVotesRepo{
		db: db,
	}
}

const (
	GetUpVoteCountQuery = "SELECT COUNT(v.vote_id) FROM votes v, pages p, sites s " +
		" WHERE v.page_id = p.page_id AND p.site_id = s.site_id AND p.site_id = ? AND p.page_id = ? "
)

func (r *upVotesRepo) GetUpVoteCounts(siteId string, pageId string) (count int, err error) {
	r.db.Raw(GetUpVoteCountQuery, siteId, pageId).Scan(&count)

	//r.db.Table("votes").Where("page_id = ?", pageId).Count(&count)

	return
}
