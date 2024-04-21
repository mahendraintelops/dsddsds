package daos

import (
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/mahendraintelops/dsddsds/dddds/pkg/rest/server/daos/clients/sqls"
	"github.com/mahendraintelops/dsddsds/dddds/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
)

type DadadDao struct {
	sqlClient *sqls.MySQLClient
}

func migrateDadads(r *sqls.MySQLClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS dadads(
		ID int NOT NULL AUTO_INCREMENT,
        
		Daaadad INT NOT NULL,
	    PRIMARY KEY (ID)
	);
	`
	_, err := r.DB.Exec(query)
	return err
}

func NewDadadDao() (*DadadDao, error) {
	sqlClient, err := sqls.InitMySQLDB()
	if err != nil {
		return nil, err
	}
	err = migrateDadads(sqlClient)
	if err != nil {
		return nil, err
	}
	return &DadadDao{
		sqlClient,
	}, nil
}

func (dadadDao *DadadDao) CreateDadad(m *models.Dadad) (*models.Dadad, error) {
	insertQuery := "INSERT INTO dadads(Daaadad) values(?)"
	res, err := dadadDao.sqlClient.DB.Exec(insertQuery, m.Daaadad)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			if mysqlErr.Number == 1062 {
				return nil, sqls.ErrDuplicate
			}
		}
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	m.Id = id
	log.Debugf("dadad created")
	return m, nil
}

func (dadadDao *DadadDao) GetDadad(id int64) (*models.Dadad, error) {
	selectQuery := "SELECT * FROM dadads WHERE Id = ?"
	row := dadadDao.sqlClient.DB.QueryRow(selectQuery, id)

	m := models.Dadad{}
	if err := row.Scan(&m.Id, &m.Daaadad); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}
	log.Debugf("dadad retrieved")
	return &m, nil
}
