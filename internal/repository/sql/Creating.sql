create database EmployeesService;


create table Employees(
                          id serial primary key not null,
                          name varchar(30),
                          age int,
                          department_name varchar(30)
)