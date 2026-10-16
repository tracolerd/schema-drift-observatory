Write-Host "Formatting Go code..." -ForegroundColor Cyan
go fmt ./...
Write-Host "`nRunning Go Vet (Linter)..." -ForegroundColor Cyan
go vet ./...
Write-Host "`nBuilding Executable..." -ForegroundColor Cyan
go build -o schema-drift-observatory.exe main.go
if ($LASTEXITCODE -eq 0) {
    Write-Host "`nBuild Successful! Executable created: schema-drift-observatory.exe" -ForegroundColor Green
} else {
    Write-Host "`nBuild Failed!" -ForegroundColor Red
}
