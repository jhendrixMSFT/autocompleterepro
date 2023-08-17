# Repro for autocomplete bug

Auto-complete should have generated the following

```go
Get: func(ctx context.Context, resourceGroupName, vmName string, options *armcompute.VirtualMachinesClientGetOptions) (resp azfake.Responder[armcompute.VirtualMachinesClientGetResponse], errResp azfake.ErrorResponder) {
},
```

Instead it generated this

```go
Get: func(ctx context.Context, resourceGroupName, vmName string, options *armcompute.VirtualMachinesClientGetOptions) (resp fake.Responder[armcompute.VirtualMachinesClientGetResponse], errResp fake.ErrorResponder) {
},
```

Note that the package name for `resp` and `errResp` return params are incorrect.  Should be `azfake` but are `fake`.
