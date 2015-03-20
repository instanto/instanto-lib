package instantolib

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Publication struct {
	Id                       int64  `json:"id"`
	Title                    string `json:"title"`
	Year                     int64  `json:"year"`
	BookTitle                string `json:"book_title"`
	Chapter                  string `json:"chapter"`
	City                     string `json:"city"`
	Country                  string `json:"country"`
	ConferenceName           string `json:"conference_name"`
	Edition                  string `json:"edition"`
	Institution              string `json:"institution"`
	Isbn                     string `json:"isbn"`
	Issn                     string `json:"issn"`
	Journal                  string `json:"journal"`
	Nationality              string `json:"nationality"`
	Number                   string `json:"number"`
	Organization             string `json:"organization"`
	Pages                    string `json:"pages"`
	School                   string `json:"school"`
	Series                   string `json:"series"`
	Volume                   string `json:"volume"`
	Language                 string `json:"language"`
	CreatedBy                string `json:"created_by"`
	UpdatedBy                string `json:"updated_by"`
	CreatedAt                int64  `json:"created_at"`
	UpdatedAt                int64  `json:"updated_at"`
	PublicationType          int64  `json:"publication_type"`
	Publisher                int64  `json:"publisher"`
	PrimaryAuthor            int64  `json:"primary_author"`
	RelMemberCreatedBy       string `json:"member_created_by,omitempty"`
	RelMemberCreatedAt       int64  `json:"member_created_at,omitempty"`
	RelResearchLineCreatedBy string `json:"research_line_created_by,omitempty"`
	RelResearchLineCreatedAt int64  `json:"research_line_created_at,omitempty"`
}

func (dbp *DBProvider) PublicationCreate(title string, year int64, bookTitle, chapter, city, country, conferenceName, edition, institution, isbn, issn, journal, language, nationality, number, organization, pages, school, series, volume, createdBy string, publicationType, publisher, primaryAuthor int64) (id int64, verr *ValidationError, err error) {
	verr = publicationValidate(title, year, bookTitle, chapter, city, country, conferenceName, edition, institution, isbn, issn, journal, language)
	if verr != nil {
		return
	}
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO publication(title,year,book_title,chapter,city,country,conference_name,edition,institution,isbn,issn,journal,language,nationality,number,organization,pages,school,series,volume,created_by,updated_by,created_at,updated_at,publication_type,publisher,primary_author) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	result, err := stmt.Exec(title, year, bookTitle, chapter, city, country, conferenceName, edition, institution, isbn, issn, journal, language, nationality, number, organization, pages, school, series, volume, createdBy, createdBy, ts, ts, publicationType, publisher, primaryAuthor)
	if err != nil {
		if is, field := IsDbError1452(err); is {
			verr = &ValidationError{field, "not exist"}
			err = nil
			return
		}
		return
	}
	id, err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}
func (dbp *DBProvider) PublicationUpdate(id int64, title string, year int64, booktitle, chapter, city, country, conferenceName, edition, institution, isbn, issn, journal, language, nationality, number, organization, pages, school, series, volume, updatedBy string, publicationType, publisher, primaryAuthor int64) (numRows int64, verr *ValidationError, err error) {
	verr = publicationValidate(title, year, booktitle, chapter, city, country, conferenceName, edition, institution, isbn, issn, journal, language)
	if verr != nil {
		return
	}
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "UPDATE publication SET title=?,year=?,book_title=?,chapter=?,city=?,country=?,conference_name=?,edition=?,institution=?,isbn=?,issn=?,journal=?,language=?,nationality=?,number=?,organization=?,pages=?,school=?,series=?,volume=?,updated_by=?,updated_at=?,publication_type=?,publisher=?,primary_author=? WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	result, err := stmt.Exec(title, year, booktitle, chapter, city, country, conferenceName, edition, institution, isbn, issn, journal, language, nationality, number, organization, pages, school, series, volume, updatedBy, ts, publicationType, publisher, primaryAuthor, id)
	if err != nil {
		if is, field := IsDbError1452(err); is {
			verr = &ValidationError{field, "not exists"}
			err = nil
			return
		}
		return
	}
	numRows, err = result.RowsAffected()
	if err != nil {
		return
	}
	return
}
func (dbp *DBProvider) PublicationDelete(id int64) (numRows int64, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM publication WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(id)
	if err != nil {
		return
	}
	numRows, err = result.RowsAffected()
	if err != nil {
		return
	}
	return
}
func (dbp *DBProvider) PublicationGetAll() (publications []*Publication, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM publication"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := Publication{}
		err = rows.Scan(&p.Id, &p.Title, &p.Year, &p.BookTitle, &p.City, &p.Chapter, &p.Country, &p.ConferenceName, &p.Edition, &p.Institution, &p.Isbn, &p.Issn, &p.Journal, &p.Language, &p.Nationality, &p.Number, &p.Organization, &p.Pages, &p.School, &p.Series, &p.Volume, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PublicationType, &p.Publisher, &p.PrimaryAuthor)
		if err != nil {
			return
		}
		publications = append(publications, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func (dbp *DBProvider) PublicationGetById(id int64) (publication *Publication, err error) {
	publication = &Publication{}
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM publication WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&publication.Id, &publication.Title, &publication.Year, &publication.BookTitle, &publication.City, &publication.Chapter, &publication.Country, &publication.ConferenceName, &publication.Edition, &publication.Institution, &publication.Isbn, &publication.Issn, &publication.Journal, &publication.Language, &publication.Nationality, &publication.Number, &publication.Organization, &publication.Pages, &publication.School, &publication.Series, &publication.Volume, &publication.CreatedBy, &publication.UpdatedBy, &publication.CreatedAt, &publication.UpdatedAt, &publication.PublicationType, &publication.Publisher, &publication.PrimaryAuthor)
	if err != nil {
		return
	}
	return
}
func (dbp *DBProvider) PublicationGetByPublicationType(publicationTypeId int64) (publications []*Publication, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM publication WHERE publication_type=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(publicationTypeId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := Publication{}
		err = rows.Scan(&p.Id, &p.Title, &p.Year, &p.BookTitle, &p.City, &p.Chapter, &p.Country, &p.ConferenceName, &p.Edition, &p.Institution, &p.Isbn, &p.Issn, &p.Journal, &p.Language, &p.Nationality, &p.Number, &p.Organization, &p.Pages, &p.School, &p.Series, &p.Volume, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PublicationType, &p.Publisher, &p.PrimaryAuthor)
		if err != nil {
			return
		}
		publications = append(publications, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func (dbp *DBProvider) PublicationGetByPublisher(publisherId int64) (publications []*Publication, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM publication WHERE publisher=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(publisherId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := Publication{}
		err = rows.Scan(&p.Id, &p.Title, &p.Year, &p.BookTitle, &p.City, &p.Chapter, &p.Country, &p.ConferenceName, &p.Edition, &p.Institution, &p.Isbn, &p.Issn, &p.Journal, &p.Language, &p.Nationality, &p.Number, &p.Organization, &p.Pages, &p.School, &p.Series, &p.Volume, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PublicationType, &p.Publisher, &p.PrimaryAuthor)
		if err != nil {
			return
		}
		publications = append(publications, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func (dbp *DBProvider) PublicationGetByPrimaryAuthor(authorId int64) (publications []*Publication, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM publication WHERE primary_author=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(authorId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := Publication{}
		err = rows.Scan(&p.Id, &p.Title, &p.Year, &p.BookTitle, &p.City, &p.Chapter, &p.Country, &p.ConferenceName, &p.Edition, &p.Institution, &p.Isbn, &p.Issn, &p.Journal, &p.Language, &p.Nationality, &p.Number, &p.Organization, &p.Pages, &p.School, &p.Series, &p.Volume, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PublicationType, &p.Publisher, &p.PrimaryAuthor)
		if err != nil {
			return
		}
		publications = append(publications, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func (dbp *DBProvider) PublicationGetByMember(memberId int64) (publications []*Publication, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT publication.*,member_publication.created_by,member_publication.created_at FROM member_publication INNER JOIN publication ON member_publication.publication=publication.id  WHERE member=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(memberId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := Publication{}
		err = rows.Scan(&p.Id, &p.Title, &p.Year, &p.BookTitle, &p.City, &p.Chapter, &p.Country, &p.ConferenceName, &p.Edition, &p.Institution, &p.Isbn, &p.Issn, &p.Journal, &p.Language, &p.Nationality, &p.Number, &p.Organization, &p.Pages, &p.School, &p.Series, &p.Volume, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PublicationType, &p.Publisher, &p.PrimaryAuthor, &p.RelMemberCreatedBy, &p.RelMemberCreatedAt)
		if err != nil {
			return
		}
		publications = append(publications, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func (dbp *DBProvider) PublicationGetByResearchLine(researchLineId int64) (publications []*Publication, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT publication.*,research_line_publication.created_by,research_line_publication.created_at FROM research_line_publication INNER JOIN publication ON research_line_publication.publication=publication.id  WHERE research_line=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(researchLineId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := Publication{}
		err = rows.Scan(&p.Id, &p.Title, &p.Year, &p.BookTitle, &p.City, &p.Chapter, &p.Country, &p.ConferenceName, &p.Edition, &p.Institution, &p.Isbn, &p.Issn, &p.Journal, &p.Language, &p.Nationality, &p.Number, &p.Organization, &p.Pages, &p.School, &p.Series, &p.Volume, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PublicationType, &p.Publisher, &p.PrimaryAuthor, &p.RelResearchLineCreatedBy, &p.RelResearchLineCreatedAt)
		if err != nil {
			return
		}
		publications = append(publications, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}

func (dbp *DBProvider) PublicationCount() (count int64, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM publication"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow().Scan(&count)
	if err != nil {
		return
	}
	return
}
func (dbp *DBProvider) PublicationExists(id int64) (exists bool, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM publication WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	var count int64
	err = stmt.QueryRow(id).Scan(&count)
	if err != nil {
		return
	}
	if count != 1 {
		return
	}
	exists = true
	return
}
func (dbp *DBProvider) PublicationAddMember(id, memberId int64, createdBy string) (verr *ValidationError, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO member_publication(member,publication,created_by,created_at) VALUES(?,?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	_, err = stmt.Exec(memberId, id, createdBy, ts)
	if err != nil {
		if IsDbError1062(err) {
			verr = &ValidationError{"member", "this member has already been added"}
			err = nil
			return
		}
		if is, field := IsDbError1452(err); is {
			verr = &ValidationError{field, "not exist"}
			err = nil
			return
		}
		return
	}
	return
}
func (dbp *DBProvider) PublicationRemoveMember(id, memberId int64) (removed bool, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM member_publication WHERE member=? AND publication=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(memberId, id)
	if err != nil {
		return
	}
	numRows, err := result.RowsAffected()
	if err != nil {
		return
	}
	if numRows != 0 {
		removed = true
	}
	return
}

func (dbp *DBProvider) PublicationGetMembers(id int64) (members []*Member, err error) {
	members, err = dbp.MemberGetByPublication(id)
	return
}
func (dbp *DBProvider) PublicationAddResearchLine(id, researchLineId int64, createdBy string) (verr *ValidationError, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO research_line_publication(research_line,publication,created_by,created_at) VALUES(?,?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	_, err = stmt.Exec(researchLineId, id, createdBy, ts)
	if err != nil {
		if IsDbError1062(err) {
			verr = &ValidationError{"research_line", "this research_line has already been added"}
			err = nil
			return
		}
		if is, field := IsDbError1452(err); is {
			verr = &ValidationError{field, "not exist"}
			err = nil
			return
		}
		return
	}
	return
}
func (dbp *DBProvider) PublicationRemoveResearchLine(id, researchLineId int64) (removed bool, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM research_line_publication WHERE research_line=? AND publication=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(researchLineId, id)
	if err != nil {
		return
	}
	numRows, err := result.RowsAffected()
	if err != nil {
		return
	}
	if numRows != 0 {
		removed = true
	}
	return
}

func (dbp *DBProvider) PublicationGetResearchLines(id int64) (researchLines []*ResearchLine, err error) {
	researchLines, err = dbp.ResearchLineGetByPublication(id)
	return
}
func (dbp *DBProvider) PublicationGetColumns() []string {
	columns := []string{
		"id",
		"title",
		"web",
		"date",
		"created_by",
		"updated_by",
		"created_at",
		"updated_at",
		"newspaper",
	}
	return columns
}
func publicationValidateTitle(title string) (verr *ValidationError) {
	if verr = validateNotEmpty("title", title); verr != nil {
		return verr
	}
	return validateLength("title", title, 200)
}
func publicationValidateYear(year int64) (err *ValidationError) {
	return validateIsNumber("year", year)
}
func publicationValidateBooktitle(booktitle string) (verr *ValidationError) {
	return validateLength("booktitle", booktitle, 200)
}
func publicationValidateChapter(chapter string) (verr *ValidationError) {
	return validateLength("chapter", chapter, 200)
}
func publicationValidateCity(city string) (verr *ValidationError) {
	return validateLength("city", city, 200)
}
func publicationValidateCountry(country string) (verr *ValidationError) {
	return validateLength("country", country, 200)
}
func publicationValidateConferenceName(conferenceName string) (verr *ValidationError) {
	return validateLength("conference_name", conferenceName, 200)
}
func publicationValidateEdition(edition string) (verr *ValidationError) {
	return validateLength("edition", edition, 200)
}
func publicationValidateInstitution(institution string) (verr *ValidationError) {
	return validateLength("institution", institution, 200)
}
func publicationValidateISBN(isbn string) (verr *ValidationError) {
	return validateLength("isbn", isbn, 200)
}
func publicationValidateISSN(issn string) (verr *ValidationError) {
	return validateLength("issn", issn, 200)
}
func publicationValidateJournal(journal string) (verr *ValidationError) {
	return validateLength("journal", journal, 200)
}
func publicationValidateLanguage(language string) (verr *ValidationError) {
	return validateLength("language", language, 200)
}

func publicationValidate(title string, year int64, booktitle, chapter, city, country, conferenceName, edition, institution, isbn, issn, journal, language string) (verr *ValidationError) {
	if verr = publicationValidateTitle(title); verr != nil {
		return verr
	}
	if verr = publicationValidateYear(year); verr != nil {
		return verr
	}
	if verr = publicationValidateBooktitle(booktitle); verr != nil {
		return verr
	}
	if verr = publicationValidateChapter(chapter); verr != nil {
		return verr
	}
	if verr = publicationValidateCity(city); verr != nil {
		return verr
	}
	if verr = publicationValidateCountry(country); verr != nil {
		return verr
	}
	if verr = publicationValidateConferenceName(conferenceName); verr != nil {
		return verr
	}
	if verr = publicationValidateEdition(edition); verr != nil {
		return verr
	}
	if verr = publicationValidateInstitution(institution); verr != nil {
		return verr
	}
	if verr = publicationValidateISBN(isbn); verr != nil {
		return verr
	}
	if verr = publicationValidateISSN(issn); verr != nil {
		return verr
	}
	if verr = publicationValidateJournal(journal); verr != nil {
		return verr
	}
	if verr = publicationValidateLanguage(language); verr != nil {
		return verr
	}
	return nil
}
