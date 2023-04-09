package query

const InsertEmployee string = `
insert into employees (name, lastname, patronymic, birthday) values($1, $2, $3, $4) returning id
`
const GetAllEmployees string = `select * from employees`
const GetByIdEmployee string = `
select id, name, lastname, patronymic, birthday from employees where id = $1
`
const DeleteByIdEmployee string = `delete from employees where id = $1 returning id`
const UpdateByIdEmployee string = `update employees set 
                     name = $2,
                     lastname = $3,
                     patronymic = $4,
                     birthday = $5
			where id = $1`

const InsertDepartment string = "insert into departments (name) values($1) returning id"
const GetByIdDepartment string = `select * from departments where id = $1`
const DeleteByIdDepartment string = `delete from departments where id = $1`
const UpdateByIdDepartment string = `update departments set name = $2 where id = $1 returning id`

const InsertReport string = "insert into reports (number, employee_id, type) values($1, $2, $3) returning id"
const GetByIdReport string = `select * from reports where id = $1`
const GetByEmployeeIdReport string = `select * from reports where employee_id = $1`
const DeleteByIdReport string = `delete from reports where id = $1 returning id`
const InsertFile string = "insert into files (filename, insert_date, data) values($1, $2, $3) returning id"
const GetByIdFile string = `select id, filename, insert_date, data from files where id = $1`
const DeleteByIdFile string = `delete from files where id = $1`
