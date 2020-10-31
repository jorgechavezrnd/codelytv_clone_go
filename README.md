<p align="center">
  <a href="http://codely.tv">
    <img src="http://codely.tv/wp-content/uploads/2016/05/cropped-logo-codelyTV.png" width="192px" height="192px"/>
  </a>
</p>

<h1 align="center">
  🐘🎯 Clone of codelytv pro platform applying Domain-Driven Design, Hexagonal Architecture and CQRS in a Monorepository
</h1>

### Commands for mooc backend application
```
go run apps/starter.go mooc_backend fake
go run apps/starter.go mooc_backend another_fake
go run apps/starter.go mooc_backend server
```

### Commands for mooc frontend application
```
go run apps/starter.go mooc_frontend fake
go run apps/starter.go mooc_frontend another_fake
go run apps/starter.go mooc_frontend server
```

### Command for run acceptance tests of mooc backend application
- Windows (Powershell)
```
cd tests/apps; godog mooc/backend; cd ../..
```

- Mac / Ubuntu
```
cd tests/apps && godog mooc/backend && cd ../..
```
