package project

import (
    "net/url"
)

type Project struct {
    Project       string
    Platform      string
    Productowner  string
    Excludeothers bool
}


func (this Project) GetQuery(platforms []string) string {
    var sqlQuery string
    if len(this.Platform) > 0 {
        sqlQuery = "Platform = \"" + this.Platform + "\""
    }
    if len(this.Project) > 0 && len(this.Platform) > 0 {
        sqlQuery = sqlQuery + " OR "
    }
    if len(this.Project) > 0 {
        sqlQuery = sqlQuery + "project = \"" + this.Project + "\""
    }

    sqlQuery = "(" + sqlQuery + ")"

    if platforms != nil {
        notInPart := " AND Platform not in ("
        for i := range platforms {
            if len(platforms[i]) > 0 {
                if i > 0 {
                    notInPart = notInPart + ","
                }
                notInPart = notInPart + "\"" + platforms[i] + "\""
            }
        }
        notInPart = notInPart + ")"
        sqlQuery = sqlQuery + notInPart
    }

    return url.QueryEscape(sqlQuery)
}
