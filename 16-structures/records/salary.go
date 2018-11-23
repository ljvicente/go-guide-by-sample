package records

type Salary struct {
    Amount int // will be exported
    Currency string // will be exported
    points int // will NOT be exported because started with lower case
}

