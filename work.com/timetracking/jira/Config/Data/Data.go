package date

type JiraData struct {
    Username string
    Password string
    Url      string
}

// TODO: those methods must be moved to somewhere they are really needed! perhaps own class together with JiraQueryGenerator
const reportName string = "ConfigureReport.jspa?"
const startDate string = "startDateId="
const endDate string = "&endDateId="
const prjId string = "&projectId="
const query string = "&jqlQueryId="
const selectedPrjId string = "&selectedProjectId="
const reportKey string = "&reportKey=com.synergyapps.plugins.jira.timepo-timesheet-plugin%3Aissues-report&Next=Next"

func (this JiraData) GetReportName() string {
    return reportName
}

func (this JiraData) GetStartDate() string {
    return startDate
}

func (this JiraData) GetEndDate() string {
    return endDate
}

func (this JiraData) GetPrjId() string {
    return prjId
}

func (this JiraData) GetSelectedPrjId() string {
    return selectedPrjId
}

func (this JiraData) GetReportKey() string {
    return reportKey
}

func (this JiraData) GetQuery() string {
    return query
}

