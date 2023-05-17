## Step
```
$go work init
$mkdir customer payment

$cd customer
$go mod init customer

$cd payment
$go mod init payment
```

Add module to workspace
```
$go work use ./customer
$go work use ./payment
```