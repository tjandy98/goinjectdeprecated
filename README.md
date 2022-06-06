# GoInject


## Introduction

**GoInject** is a dynamic-link library injection program written in Go. It injects dll into a target process.
## Information

Current injection technique utilises `CreateRemoteThread` and `LoadLibrary` from the Win32 API.


