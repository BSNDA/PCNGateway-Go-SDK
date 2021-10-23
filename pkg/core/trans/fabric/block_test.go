package fabric

import (
	"fmt"
	"gotest.tools/assert"
	"testing"
)

func TestConvertToBlock(t *testing.T) {
	data := `CkYIARIgp5sq0cjaMveBt18j6v8AOhXW4AaE8aiJVXAnOd0FPK4aIFFcnJcGilTIF6VLTckmkJhrGY+wgCwuF0qF7oF/lIDsEpMkCpAkCsUjCvcJCnoIAxoMCI2XkIoGEOGw3pUBIhxhcHAwMDAxMjAyMTA5MTcxMTI1NTYyNDM1NzYwKkBjYmNkMGZlNjU3YzQ4MzY4NjkxNjk2ZDg3NjM2ZmM4NGVkMjc2MTE2YzE4YWIyOTE0OWQyMjNmNTZiZGUyMjVhOggSBhIEbHNjYxL4CArbCAoLT3JnYk5vZGVNU1ASywgtLS0tLUJFR0lOIENFUlRJRklDQVRFLS0tLS0KTUlJQy9qQ0NBcVdnQXdJQkFnSVVGRjRJMk16V1NBRW1UOWFVZnMvN09YUVlnWEV3Q2dZSUtvWkl6ajBFQXdJdwpUakVMTUFrR0ExVUVCaE1DUTA0eEVEQU9CZ05WQkFnVEIwSmxhV3BwYm1jeEREQUtCZ05WQkFvVEEwSlRUakVQCk1BMEdBMVVFQ3hNR1kyeHBaVzUwTVE0d0RBWURWUVFERXdWaWMyNWpZVEFnRncweU1EQTBNRGt4TURVNU1EQmEKR0E4eU1UQXdNRE15TVRFeE1EUXdNRm93Z1pBeEN6QUpCZ05WQkFZVEFrTk9NUkF3RGdZRFZRUUlFd2RDWldscQphVzVuTVF3d0NnWURWUVFLRXdOQ1UwNHhQREFOQmdOVkJBc1RCbU5zYVdWdWREQVBCZ05WQkFzVENHOXlaMkp1CmIyUmxNQTRHQTFVRUN4TUhZbk51WW1GelpUQUtCZ05WQkFzVEEyTnZiVEVqTUNFR0ExVUVBd3dhUVdSdGFXNUEKYjNKblltNXZaR1V1WW5OdVltRnpaUzVqYjIwd1dUQVRCZ2NxaGtqT1BRSUJCZ2dxaGtqT1BRTUJCd05DQUFTOQoyanY5QzV0eWhaV1VHM0Jmb3RwQWNvK1duUmdWbVdUeGpiTkM3akJGb2twV3RlWUdYY29yRExEcXAvd3dGWmlPCjlXN1ZWeEV2ZXdHbG0xQ1BmT0d0bzRJQkdqQ0NBUll3RGdZRFZSMFBBUUgvQkFRREFnZUFNQXdHQTFVZEV3RUIKL3dRQ01BQXdIUVlEVlIwT0JCWUVGRFd5Zk0rZG53RmVvT1BlTThhZUs1NUNXdFhiTUI4R0ExVWRJd1FZTUJhQQpGQWNJNEgra0lzOHZuOTRaWVlwa3JkKzVsZE1LTUNJR0ExVWRFUVFiTUJtQ0YyTmhMbTl5WjJKdWIyUmxMbUp6CmJtSmhjMlV1WTI5dE1JR1JCZ2dxQXdRRkJnY0lBUVNCaEhzaVlYUjBjbk1pT25zaWFHWXVRV1ptYVd4cFlYUnAKYjI0aU9pSnZjbWRpYm05a1pTNWljMjVpWVhObExtTnZiU0lzSW1obUxrVnVjbTlzYkcxbGJuUkpSQ0k2SWtGawpiV2x1UUc5eVoySnViMlJsTG1KemJtSmhjMlV1WTI5dElpd2lhR1l1Vkhsd1pTSTZJbU5zYVdWdWRDSXNJbkp2CmJHVWlPaUpoWkcxcGJpSjlmVEFLQmdncWhrak9QUVFEQWdOSEFEQkVBaUJEVlhramw5Ukh6RlBvb0kzbWk1QmMKSUxJVGlaWTNHQ2xqaDBvMXBEYTJWQUlnQVBUR1pkanNVYkRSY25CVVEvN0JzVzlPemV1VU1iYW1sWVl2Z2hTRgpNblk9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0KEhgHJDn/FqbbHH2zZbZOJsAsMpxiyqgACScSyBkKxRkK+AgK2wgKC09yZ2JOb2RlTVNQEssILS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUMvakNDQXFXZ0F3SUJBZ0lVRkY0STJNeldTQUVtVDlhVWZzLzdPWFFZZ1hFd0NnWUlLb1pJemowRUF3SXcKVGpFTE1Ba0dBMVVFQmhNQ1EwNHhFREFPQmdOVkJBZ1RCMEpsYVdwcGJtY3hEREFLQmdOVkJBb1RBMEpUVGpFUApNQTBHQTFVRUN4TUdZMnhwWlc1ME1RNHdEQVlEVlFRREV3VmljMjVqWVRBZ0Z3MHlNREEwTURreE1EVTVNREJhCkdBOHlNVEF3TURNeU1URXhNRFF3TUZvd2daQXhDekFKQmdOVkJBWVRBa05PTVJBd0RnWURWUVFJRXdkQ1pXbHEKYVc1bk1Rd3dDZ1lEVlFRS0V3TkNVMDR4UERBTkJnTlZCQXNUQm1Oc2FXVnVkREFQQmdOVkJBc1RDRzl5WjJKdQpiMlJsTUE0R0ExVUVDeE1IWW5OdVltRnpaVEFLQmdOVkJBc1RBMk52YlRFak1DRUdBMVVFQXd3YVFXUnRhVzVBCmIzSm5ZbTV2WkdVdVluTnVZbUZ6WlM1amIyMHdXVEFUQmdjcWhrak9QUUlCQmdncWhrak9QUU1CQndOQ0FBUzkKMmp2OUM1dHloWldVRzNCZm90cEFjbytXblJnVm1XVHhqYk5DN2pCRm9rcFd0ZVlHWGNvckRMRHFwL3d3RlppTwo5VzdWVnhFdmV3R2xtMUNQZk9HdG80SUJHakNDQVJZd0RnWURWUjBQQVFIL0JBUURBZ2VBTUF3R0ExVWRFd0VCCi93UUNNQUF3SFFZRFZSME9CQllFRkRXeWZNK2Rud0Zlb09QZU04YWVLNTVDV3RYYk1COEdBMVVkSXdRWU1CYUEKRkFjSTRIK2tJczh2bjk0WllZcGtyZCs1bGRNS01DSUdBMVVkRVFRYk1CbUNGMk5oTG05eVoySnViMlJsTG1KegpibUpoYzJVdVkyOXRNSUdSQmdncUF3UUZCZ2NJQVFTQmhIc2lZWFIwY25NaU9uc2lhR1l1UVdabWFXeHBZWFJwCmIyNGlPaUp2Y21kaWJtOWtaUzVpYzI1aVlYTmxMbU52YlNJc0ltaG1Ma1Z1Y205c2JHMWxiblJKUkNJNklrRmsKYldsdVFHOXlaMkp1YjJSbExtSnpibUpoYzJVdVkyOXRJaXdpYUdZdVZIbHdaU0k2SW1Oc2FXVnVkQ0lzSW5KdgpiR1VpT2lKaFpHMXBiaUo5ZlRBS0JnZ3Foa2pPUFFRREFnTkhBREJFQWlCRFZYa2psOVJIekZQb29JM21pNUJjCklMSVRpWlkzR0NsamgwbzFwRGEyVkFJZ0FQVEdaZGpzVWJEUmNuQlVRLzdCc1c5T3pldVVNYmFtbFlZdmdoU0YKTW5ZPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tChIYByQ5/xam2xx9s2W2TibALDKcYsqoAAknEscQCo4BCosBCogBCAESBhIEbHNjYxp8CgZkZXBsb3kKHGFwcDAwMDEyMDIxMDkxNzExMjU1NjI0MzU3NjAKRgpECAESOAoJYnNuQmFzZUNDEiJjY19hcHAwMDAxMjAyMTA5MTcxMTI1NTYyNDM1NzYwXzAxGgcxLjAuMC4xGgYKBGluaXQKAAoEZXNjYwoEdnNjYxKzDwqfBgogTD8IVGFJR8cJ+g+pZZ1AQNIBIDJ/ce8XPhgnJoFGHvYS+gUK1gMSbwoiY2NfYXBwMDAwMTIwMjEwOTE3MTEyNTU2MjQzNTc2MF8wMRJJGkcKB2NjX2tleV8aPHsiQmFzZUtleSI6ImNjX2tleV8iLCJCYXNlSW5mbyI6IldlbGNvbWUgdG8gdXNlIENoYWluQ29kZSAifRLiAgoEbHNjYxLZAgokCiJjY19hcHAwMDAxMjAyMTA5MTcxMTI1NTYyNDM1NzYwXzAxGrACCiJjY19hcHAwMDAxMjAyMTA5MTcxMTI1NTYyNDM1NzYwXzAxGokCCiJjY19hcHAwMDAxMjAyMTA5MTcxMTI1NTYyNDM1NzYwXzAxEgcxLjAuMC4xGgRlc2NjIgR2c2NjKjASDBIKCAESAggAEgIIARoPEg0KC09yZ2FOb2RlTVNQGg8SDQoLT3JnYk5vZGVNU1AyRAogNeL1N4XA/D/WDEEpCCYA+1R2B6iTtuenSoCXxT+ffHESINaBDL73hWXawE774jM7RvlqFgdRW62z2tvPTDI/MbYUOiB+y7m9zanke7YvGJ2cG9zxmK1Mrq9BpIUZVt6hlHw3J0I0EgwSCggBEgIIABICCAEaERIPCgtPcmdhTm9kZU1TUBABGhESDwoLT3JnYk5vZGVNU1AQARqPAgjIARqJAgoiY2NfYXBwMDAwMTIwMjEwOTE3MTEyNTU2MjQzNTc2MF8wMRIHMS4wLjAuMRoEZXNjYyIEdnNjYyowEgwSCggBEgIIABICCAEaDxINCgtPcmdhTm9kZU1TUBoPEg0KC09yZ2JOb2RlTVNQMkQKIDXi9TeFwPw/1gxBKQgmAPtUdgeok7bnp0qAl8U/n3xxEiDWgQy+94Vl2sBO++IzO0b5ahYHUVuts9rbz0wyPzG2FDogfsu5vc2p5Hu2LxidnBvc8ZitTK6vQaSFGVbeoZR8NydCNBIMEgoIARICCAASAggBGhESDwoLT3JnYU5vZGVNU1AQARoREg8KC09yZ2JOb2RlTVNQEAEiDRIEbHNjYxoFMS40LjMSjgkKwggKC09yZ2JOb2RlTVNQErIILS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUM3RENDQXBPZ0F3SUJBZ0lVQ2VEYzFpalZtZDMwRHR0RGJNYTlsY0FXV1ljd0NnWUlLb1pJemowRUF3SXcKVGpFTE1Ba0dBMVVFQmhNQ1EwNHhFREFPQmdOVkJBZ1RCMEpsYVdwcGJtY3hEREFLQmdOVkJBb1RBMEpUVGpFUApNQTBHQTFVRUN4TUdZMnhwWlc1ME1RNHdEQVlEVlFRREV3VmljMjVqWVRBZ0Z3MHlNREEwTURreE1EVTVNREJhCkdBOHlNVEF3TURNeU1URXhNRFF3TUZvd2dZNHhDekFKQmdOVkJBWVRBa05PTVJBd0RnWURWUVFJRXdkQ1pXbHEKYVc1bk1Rd3dDZ1lEVlFRS0V3TkNVMDR4T2pBTEJnTlZCQXNUQkhCbFpYSXdEd1lEVlFRTEV3aHZjbWRpYm05awpaVEFPQmdOVkJBc1RCMkp6Ym1KaGMyVXdDZ1lEVlFRTEV3TmpiMjB4SXpBaEJnTlZCQU1UR25CbFpYSXhMbTl5CloySnViMlJsTG1KemJtSmhjMlV1WTI5dE1Ga3dFd1lIS29aSXpqMENBUVlJS29aSXpqMERBUWNEUWdBRVRwS2MKZnlHSFJGaDZVcDJnSUhTVldGZDloT3FkaUZxUUxPNkJkVzBsdStHZTgxdlg1Y3Fka1NsRG53Z1ZsZm5WRU1RdwptVWJxRzN0SW1xdTcyOE8vWnFPQ0FRb3dnZ0VHTUE0R0ExVWREd0VCL3dRRUF3SUhnREFNQmdOVkhSTUJBZjhFCkFqQUFNQjBHQTFVZERnUVdCQlJFL2M1ekNLaG9zQTQ0ZHJ1bkpucWxvdllnTURBZkJnTlZIU01FR0RBV2dCUUgKQ09CL3BDTFBMNS9lR1dHS1pLM2Z1WlhUQ2pBbEJnTlZIUkVFSGpBY2docHdaV1Z5TVM1dmNtZGlibTlrWlM1aQpjMjVpWVhObExtTnZiVEIvQmdncUF3UUZCZ2NJQVFSemV5SmhkSFJ5Y3lJNmV5Sm9aaTVCWm1acGJHbGhkR2x2CmJpSTZJbTl5WjJKdWIyUmxMbUp6Ym1KaGMyVXVZMjl0SWl3aWFHWXVSVzV5YjJ4c2JXVnVkRWxFSWpvaWNHVmwKY2pFdWIzSm5ZbTV2WkdVdVluTnVZbUZ6WlM1amIyMGlMQ0pvWmk1VWVYQmxJam9pY0dWbGNpSjlmVEFLQmdncQpoa2pPUFFRREFnTkhBREJFQWlCay9jeUx5dnhOdVBuNlFJYlhlc2luYmgwbCt0a2dqVmVBcEd1YjA2MHIwUUlnCmJQcm5JUUtaYzNLSkpvSlBLQVQvSWh5NGlaUmN3dlR6bVFmOGFRWFhMaG89Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0KEkcwRQIhAOJsMVVxsH9PucLu5xgJEZ0kPR7roBEHOipOlNfsTTr3AiB2VaYme4iuMr/5rg2ZS4TrD6gbSRBpP0k2w9emXgYxUxJGMEQCIDt4Rbmex/bL6D4ayUioBdkJ+Wz8GSNXW+O0CtZrZZTDAiB8odh3piPYlgSuvs9jpeEsIaZnS/9zMkFhcIF6FmdYSRqLCgrSCQoPCgASCwoJCgMBAgMQBBgFEr4JCvIICtUICgpPcmRlcmVyTVNQEsYILS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUMvRENDQXFLZ0F3SUJBZ0lVY0hyOVRvYVV1WmF2USsweW9Qai90aTRGQkhjd0NnWUlLb1pJemowRUF3SXcKVGpFTE1Ba0dBMVVFQmhNQ1EwNHhFREFPQmdOVkJBZ1RCMEpsYVdwcGJtY3hEREFLQmdOVkJBb1RBMEpUVGpFUApNQTBHQTFVRUN4TUdZMnhwWlc1ME1RNHdEQVlEVlFRREV3VmljMjVqWVRBZ0Z3MHlNREEwTURrd09USXlNREJhCkdBOHlNVEF3TURNeU1UQTVNVGt3TUZvd2daUXhDekFKQmdOVkJBWVRBa05PTVJBd0RnWURWUVFJRXdkQ1pXbHEKYVc1bk1Rd3dDZ1lEVlFRS0V3TkNVMDR4UGpBT0JnTlZCQXNUQjI5eVpHVnlaWEl3RUFZRFZRUUxFd2x2Y21SbApjbTV2WkdVd0RnWURWUVFMRXdkaWMyNWlZWE5sTUFvR0ExVUVDeE1EWTI5dE1TVXdJd1lEVlFRREV4eHZjbVJsCmNqSXViM0prWlhKdWIyUmxMbUp6Ym1KaGMyVXVZMjl0TUZrd0V3WUhLb1pJemowQ0FRWUlLb1pJemowREFRY0QKUWdBRVZmSXV0Nk9IQUpJSVA2ZTFVUVpEQzBTSFJnZGhMOHhDbmNYN1cvQUtNb01yQVF3SmFHMDl0VWNZVTROWApTeFdkbHB4QXJDcUhrUUtETFp6eE90RWIrS09DQVJNd2dnRVBNQTRHQTFVZER3RUIvd1FFQXdJSGdEQU1CZ05WCkhSTUJBZjhFQWpBQU1CMEdBMVVkRGdRV0JCVGxvM24rcDlNeUJwRGJZdTR2ZXdqbklRZWpOREFmQmdOVkhTTUUKR0RBV2dCVGx1UStzelFsNDc5cE5qSXZWRVlsMEFqWCtVekFuQmdOVkhSRUVJREFlZ2h4dmNtUmxjakl1YjNKawpaWEp1YjJSbExtSnpibUpoYzJVdVkyOXRNSUdGQmdncUF3UUZCZ2NJQVFSNWV5SmhkSFJ5Y3lJNmV5Sm9aaTVCClptWnBiR2xoZEdsdmJpSTZJbTl5WkdWeWJtOWtaUzVpYzI1aVlYTmxMbU52YlNJc0ltaG1Ma1Z1Y205c2JHMWwKYm5SSlJDSTZJbTl5WkdWeU1pNXZjbVJsY201dlpHVXVZbk51WW1GelpTNWpiMjBpTENKb1ppNVVlWEJsSWpvaQpiM0prWlhKbGNpSjlmVEFLQmdncWhrak9QUVFEQWdOSUFEQkZBaUVBMUNheEFIQ05yUzhvN1E2VVZFR0xvSXZMCnhuSE1jaEJ4SjlZb2RPM0x1R0VDSUhRWWtXa1FRcjhWaWFHdXFFYmFXdm0zdkpVTEcyYW14WXd4VVZDUW16ekMKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQoSGGCs3OHkMgkSDyDeSROyIEut5dLCDeV5lBJHMEUCIQCkIDJwx0RpnOtn7MP0fINshaDy3jBbY7J8ItKGyuN2NwIgR0AVm6dRF7tC4Td2GzTUrtZFbbE3nuvOadLELeNEpXAKAAoBAAoLCgkKAwECAxAEGAUKIgog8crPDCZCfI46g/Qwl9tMVviJU8AJjnNsyk+fwFOPUIM=`

	block, err := ConvertToBlock(data)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, block.Header.Number, uint64(1))
}

func TestConvertBlockToJson(t *testing.T) {
	data := `CkYIARIgp5sq0cjaMveBt18j6v8AOhXW4AaE8aiJVXAnOd0FPK4aIFFcnJcGilTIF6VLTckmkJhrGY+wgCwuF0qF7oF/lIDsEpMkCpAkCsUjCvcJCnoIAxoMCI2XkIoGEOGw3pUBIhxhcHAwMDAxMjAyMTA5MTcxMTI1NTYyNDM1NzYwKkBjYmNkMGZlNjU3YzQ4MzY4NjkxNjk2ZDg3NjM2ZmM4NGVkMjc2MTE2YzE4YWIyOTE0OWQyMjNmNTZiZGUyMjVhOggSBhIEbHNjYxL4CArbCAoLT3JnYk5vZGVNU1ASywgtLS0tLUJFR0lOIENFUlRJRklDQVRFLS0tLS0KTUlJQy9qQ0NBcVdnQXdJQkFnSVVGRjRJMk16V1NBRW1UOWFVZnMvN09YUVlnWEV3Q2dZSUtvWkl6ajBFQXdJdwpUakVMTUFrR0ExVUVCaE1DUTA0eEVEQU9CZ05WQkFnVEIwSmxhV3BwYm1jeEREQUtCZ05WQkFvVEEwSlRUakVQCk1BMEdBMVVFQ3hNR1kyeHBaVzUwTVE0d0RBWURWUVFERXdWaWMyNWpZVEFnRncweU1EQTBNRGt4TURVNU1EQmEKR0E4eU1UQXdNRE15TVRFeE1EUXdNRm93Z1pBeEN6QUpCZ05WQkFZVEFrTk9NUkF3RGdZRFZRUUlFd2RDWldscQphVzVuTVF3d0NnWURWUVFLRXdOQ1UwNHhQREFOQmdOVkJBc1RCbU5zYVdWdWREQVBCZ05WQkFzVENHOXlaMkp1CmIyUmxNQTRHQTFVRUN4TUhZbk51WW1GelpUQUtCZ05WQkFzVEEyTnZiVEVqTUNFR0ExVUVBd3dhUVdSdGFXNUEKYjNKblltNXZaR1V1WW5OdVltRnpaUzVqYjIwd1dUQVRCZ2NxaGtqT1BRSUJCZ2dxaGtqT1BRTUJCd05DQUFTOQoyanY5QzV0eWhaV1VHM0Jmb3RwQWNvK1duUmdWbVdUeGpiTkM3akJGb2twV3RlWUdYY29yRExEcXAvd3dGWmlPCjlXN1ZWeEV2ZXdHbG0xQ1BmT0d0bzRJQkdqQ0NBUll3RGdZRFZSMFBBUUgvQkFRREFnZUFNQXdHQTFVZEV3RUIKL3dRQ01BQXdIUVlEVlIwT0JCWUVGRFd5Zk0rZG53RmVvT1BlTThhZUs1NUNXdFhiTUI4R0ExVWRJd1FZTUJhQQpGQWNJNEgra0lzOHZuOTRaWVlwa3JkKzVsZE1LTUNJR0ExVWRFUVFiTUJtQ0YyTmhMbTl5WjJKdWIyUmxMbUp6CmJtSmhjMlV1WTI5dE1JR1JCZ2dxQXdRRkJnY0lBUVNCaEhzaVlYUjBjbk1pT25zaWFHWXVRV1ptYVd4cFlYUnAKYjI0aU9pSnZjbWRpYm05a1pTNWljMjVpWVhObExtTnZiU0lzSW1obUxrVnVjbTlzYkcxbGJuUkpSQ0k2SWtGawpiV2x1UUc5eVoySnViMlJsTG1KemJtSmhjMlV1WTI5dElpd2lhR1l1Vkhsd1pTSTZJbU5zYVdWdWRDSXNJbkp2CmJHVWlPaUpoWkcxcGJpSjlmVEFLQmdncWhrak9QUVFEQWdOSEFEQkVBaUJEVlhramw5Ukh6RlBvb0kzbWk1QmMKSUxJVGlaWTNHQ2xqaDBvMXBEYTJWQUlnQVBUR1pkanNVYkRSY25CVVEvN0JzVzlPemV1VU1iYW1sWVl2Z2hTRgpNblk9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0KEhgHJDn/FqbbHH2zZbZOJsAsMpxiyqgACScSyBkKxRkK+AgK2wgKC09yZ2JOb2RlTVNQEssILS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUMvakNDQXFXZ0F3SUJBZ0lVRkY0STJNeldTQUVtVDlhVWZzLzdPWFFZZ1hFd0NnWUlLb1pJemowRUF3SXcKVGpFTE1Ba0dBMVVFQmhNQ1EwNHhFREFPQmdOVkJBZ1RCMEpsYVdwcGJtY3hEREFLQmdOVkJBb1RBMEpUVGpFUApNQTBHQTFVRUN4TUdZMnhwWlc1ME1RNHdEQVlEVlFRREV3VmljMjVqWVRBZ0Z3MHlNREEwTURreE1EVTVNREJhCkdBOHlNVEF3TURNeU1URXhNRFF3TUZvd2daQXhDekFKQmdOVkJBWVRBa05PTVJBd0RnWURWUVFJRXdkQ1pXbHEKYVc1bk1Rd3dDZ1lEVlFRS0V3TkNVMDR4UERBTkJnTlZCQXNUQm1Oc2FXVnVkREFQQmdOVkJBc1RDRzl5WjJKdQpiMlJsTUE0R0ExVUVDeE1IWW5OdVltRnpaVEFLQmdOVkJBc1RBMk52YlRFak1DRUdBMVVFQXd3YVFXUnRhVzVBCmIzSm5ZbTV2WkdVdVluTnVZbUZ6WlM1amIyMHdXVEFUQmdjcWhrak9QUUlCQmdncWhrak9QUU1CQndOQ0FBUzkKMmp2OUM1dHloWldVRzNCZm90cEFjbytXblJnVm1XVHhqYk5DN2pCRm9rcFd0ZVlHWGNvckRMRHFwL3d3RlppTwo5VzdWVnhFdmV3R2xtMUNQZk9HdG80SUJHakNDQVJZd0RnWURWUjBQQVFIL0JBUURBZ2VBTUF3R0ExVWRFd0VCCi93UUNNQUF3SFFZRFZSME9CQllFRkRXeWZNK2Rud0Zlb09QZU04YWVLNTVDV3RYYk1COEdBMVVkSXdRWU1CYUEKRkFjSTRIK2tJczh2bjk0WllZcGtyZCs1bGRNS01DSUdBMVVkRVFRYk1CbUNGMk5oTG05eVoySnViMlJsTG1KegpibUpoYzJVdVkyOXRNSUdSQmdncUF3UUZCZ2NJQVFTQmhIc2lZWFIwY25NaU9uc2lhR1l1UVdabWFXeHBZWFJwCmIyNGlPaUp2Y21kaWJtOWtaUzVpYzI1aVlYTmxMbU52YlNJc0ltaG1Ma1Z1Y205c2JHMWxiblJKUkNJNklrRmsKYldsdVFHOXlaMkp1YjJSbExtSnpibUpoYzJVdVkyOXRJaXdpYUdZdVZIbHdaU0k2SW1Oc2FXVnVkQ0lzSW5KdgpiR1VpT2lKaFpHMXBiaUo5ZlRBS0JnZ3Foa2pPUFFRREFnTkhBREJFQWlCRFZYa2psOVJIekZQb29JM21pNUJjCklMSVRpWlkzR0NsamgwbzFwRGEyVkFJZ0FQVEdaZGpzVWJEUmNuQlVRLzdCc1c5T3pldVVNYmFtbFlZdmdoU0YKTW5ZPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tChIYByQ5/xam2xx9s2W2TibALDKcYsqoAAknEscQCo4BCosBCogBCAESBhIEbHNjYxp8CgZkZXBsb3kKHGFwcDAwMDEyMDIxMDkxNzExMjU1NjI0MzU3NjAKRgpECAESOAoJYnNuQmFzZUNDEiJjY19hcHAwMDAxMjAyMTA5MTcxMTI1NTYyNDM1NzYwXzAxGgcxLjAuMC4xGgYKBGluaXQKAAoEZXNjYwoEdnNjYxKzDwqfBgogTD8IVGFJR8cJ+g+pZZ1AQNIBIDJ/ce8XPhgnJoFGHvYS+gUK1gMSbwoiY2NfYXBwMDAwMTIwMjEwOTE3MTEyNTU2MjQzNTc2MF8wMRJJGkcKB2NjX2tleV8aPHsiQmFzZUtleSI6ImNjX2tleV8iLCJCYXNlSW5mbyI6IldlbGNvbWUgdG8gdXNlIENoYWluQ29kZSAifRLiAgoEbHNjYxLZAgokCiJjY19hcHAwMDAxMjAyMTA5MTcxMTI1NTYyNDM1NzYwXzAxGrACCiJjY19hcHAwMDAxMjAyMTA5MTcxMTI1NTYyNDM1NzYwXzAxGokCCiJjY19hcHAwMDAxMjAyMTA5MTcxMTI1NTYyNDM1NzYwXzAxEgcxLjAuMC4xGgRlc2NjIgR2c2NjKjASDBIKCAESAggAEgIIARoPEg0KC09yZ2FOb2RlTVNQGg8SDQoLT3JnYk5vZGVNU1AyRAogNeL1N4XA/D/WDEEpCCYA+1R2B6iTtuenSoCXxT+ffHESINaBDL73hWXawE774jM7RvlqFgdRW62z2tvPTDI/MbYUOiB+y7m9zanke7YvGJ2cG9zxmK1Mrq9BpIUZVt6hlHw3J0I0EgwSCggBEgIIABICCAEaERIPCgtPcmdhTm9kZU1TUBABGhESDwoLT3JnYk5vZGVNU1AQARqPAgjIARqJAgoiY2NfYXBwMDAwMTIwMjEwOTE3MTEyNTU2MjQzNTc2MF8wMRIHMS4wLjAuMRoEZXNjYyIEdnNjYyowEgwSCggBEgIIABICCAEaDxINCgtPcmdhTm9kZU1TUBoPEg0KC09yZ2JOb2RlTVNQMkQKIDXi9TeFwPw/1gxBKQgmAPtUdgeok7bnp0qAl8U/n3xxEiDWgQy+94Vl2sBO++IzO0b5ahYHUVuts9rbz0wyPzG2FDogfsu5vc2p5Hu2LxidnBvc8ZitTK6vQaSFGVbeoZR8NydCNBIMEgoIARICCAASAggBGhESDwoLT3JnYU5vZGVNU1AQARoREg8KC09yZ2JOb2RlTVNQEAEiDRIEbHNjYxoFMS40LjMSjgkKwggKC09yZ2JOb2RlTVNQErIILS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUM3RENDQXBPZ0F3SUJBZ0lVQ2VEYzFpalZtZDMwRHR0RGJNYTlsY0FXV1ljd0NnWUlLb1pJemowRUF3SXcKVGpFTE1Ba0dBMVVFQmhNQ1EwNHhFREFPQmdOVkJBZ1RCMEpsYVdwcGJtY3hEREFLQmdOVkJBb1RBMEpUVGpFUApNQTBHQTFVRUN4TUdZMnhwWlc1ME1RNHdEQVlEVlFRREV3VmljMjVqWVRBZ0Z3MHlNREEwTURreE1EVTVNREJhCkdBOHlNVEF3TURNeU1URXhNRFF3TUZvd2dZNHhDekFKQmdOVkJBWVRBa05PTVJBd0RnWURWUVFJRXdkQ1pXbHEKYVc1bk1Rd3dDZ1lEVlFRS0V3TkNVMDR4T2pBTEJnTlZCQXNUQkhCbFpYSXdEd1lEVlFRTEV3aHZjbWRpYm05awpaVEFPQmdOVkJBc1RCMkp6Ym1KaGMyVXdDZ1lEVlFRTEV3TmpiMjB4SXpBaEJnTlZCQU1UR25CbFpYSXhMbTl5CloySnViMlJsTG1KemJtSmhjMlV1WTI5dE1Ga3dFd1lIS29aSXpqMENBUVlJS29aSXpqMERBUWNEUWdBRVRwS2MKZnlHSFJGaDZVcDJnSUhTVldGZDloT3FkaUZxUUxPNkJkVzBsdStHZTgxdlg1Y3Fka1NsRG53Z1ZsZm5WRU1RdwptVWJxRzN0SW1xdTcyOE8vWnFPQ0FRb3dnZ0VHTUE0R0ExVWREd0VCL3dRRUF3SUhnREFNQmdOVkhSTUJBZjhFCkFqQUFNQjBHQTFVZERnUVdCQlJFL2M1ekNLaG9zQTQ0ZHJ1bkpucWxvdllnTURBZkJnTlZIU01FR0RBV2dCUUgKQ09CL3BDTFBMNS9lR1dHS1pLM2Z1WlhUQ2pBbEJnTlZIUkVFSGpBY2docHdaV1Z5TVM1dmNtZGlibTlrWlM1aQpjMjVpWVhObExtTnZiVEIvQmdncUF3UUZCZ2NJQVFSemV5SmhkSFJ5Y3lJNmV5Sm9aaTVCWm1acGJHbGhkR2x2CmJpSTZJbTl5WjJKdWIyUmxMbUp6Ym1KaGMyVXVZMjl0SWl3aWFHWXVSVzV5YjJ4c2JXVnVkRWxFSWpvaWNHVmwKY2pFdWIzSm5ZbTV2WkdVdVluTnVZbUZ6WlM1amIyMGlMQ0pvWmk1VWVYQmxJam9pY0dWbGNpSjlmVEFLQmdncQpoa2pPUFFRREFnTkhBREJFQWlCay9jeUx5dnhOdVBuNlFJYlhlc2luYmgwbCt0a2dqVmVBcEd1YjA2MHIwUUlnCmJQcm5JUUtaYzNLSkpvSlBLQVQvSWh5NGlaUmN3dlR6bVFmOGFRWFhMaG89Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0KEkcwRQIhAOJsMVVxsH9PucLu5xgJEZ0kPR7roBEHOipOlNfsTTr3AiB2VaYme4iuMr/5rg2ZS4TrD6gbSRBpP0k2w9emXgYxUxJGMEQCIDt4Rbmex/bL6D4ayUioBdkJ+Wz8GSNXW+O0CtZrZZTDAiB8odh3piPYlgSuvs9jpeEsIaZnS/9zMkFhcIF6FmdYSRqLCgrSCQoPCgASCwoJCgMBAgMQBBgFEr4JCvIICtUICgpPcmRlcmVyTVNQEsYILS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUMvRENDQXFLZ0F3SUJBZ0lVY0hyOVRvYVV1WmF2USsweW9Qai90aTRGQkhjd0NnWUlLb1pJemowRUF3SXcKVGpFTE1Ba0dBMVVFQmhNQ1EwNHhFREFPQmdOVkJBZ1RCMEpsYVdwcGJtY3hEREFLQmdOVkJBb1RBMEpUVGpFUApNQTBHQTFVRUN4TUdZMnhwWlc1ME1RNHdEQVlEVlFRREV3VmljMjVqWVRBZ0Z3MHlNREEwTURrd09USXlNREJhCkdBOHlNVEF3TURNeU1UQTVNVGt3TUZvd2daUXhDekFKQmdOVkJBWVRBa05PTVJBd0RnWURWUVFJRXdkQ1pXbHEKYVc1bk1Rd3dDZ1lEVlFRS0V3TkNVMDR4UGpBT0JnTlZCQXNUQjI5eVpHVnlaWEl3RUFZRFZRUUxFd2x2Y21SbApjbTV2WkdVd0RnWURWUVFMRXdkaWMyNWlZWE5sTUFvR0ExVUVDeE1EWTI5dE1TVXdJd1lEVlFRREV4eHZjbVJsCmNqSXViM0prWlhKdWIyUmxMbUp6Ym1KaGMyVXVZMjl0TUZrd0V3WUhLb1pJemowQ0FRWUlLb1pJemowREFRY0QKUWdBRVZmSXV0Nk9IQUpJSVA2ZTFVUVpEQzBTSFJnZGhMOHhDbmNYN1cvQUtNb01yQVF3SmFHMDl0VWNZVTROWApTeFdkbHB4QXJDcUhrUUtETFp6eE90RWIrS09DQVJNd2dnRVBNQTRHQTFVZER3RUIvd1FFQXdJSGdEQU1CZ05WCkhSTUJBZjhFQWpBQU1CMEdBMVVkRGdRV0JCVGxvM24rcDlNeUJwRGJZdTR2ZXdqbklRZWpOREFmQmdOVkhTTUUKR0RBV2dCVGx1UStzelFsNDc5cE5qSXZWRVlsMEFqWCtVekFuQmdOVkhSRUVJREFlZ2h4dmNtUmxjakl1YjNKawpaWEp1YjJSbExtSnpibUpoYzJVdVkyOXRNSUdGQmdncUF3UUZCZ2NJQVFSNWV5SmhkSFJ5Y3lJNmV5Sm9aaTVCClptWnBiR2xoZEdsdmJpSTZJbTl5WkdWeWJtOWtaUzVpYzI1aVlYTmxMbU52YlNJc0ltaG1Ma1Z1Y205c2JHMWwKYm5SSlJDSTZJbTl5WkdWeU1pNXZjbVJsY201dlpHVXVZbk51WW1GelpTNWpiMjBpTENKb1ppNVVlWEJsSWpvaQpiM0prWlhKbGNpSjlmVEFLQmdncWhrak9QUVFEQWdOSUFEQkZBaUVBMUNheEFIQ05yUzhvN1E2VVZFR0xvSXZMCnhuSE1jaEJ4SjlZb2RPM0x1R0VDSUhRWWtXa1FRcjhWaWFHdXFFYmFXdm0zdkpVTEcyYW14WXd4VVZDUW16ekMKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQoSGGCs3OHkMgkSDyDeSROyIEut5dLCDeV5lBJHMEUCIQCkIDJwx0RpnOtn7MP0fINshaDy3jBbY7J8ItKGyuN2NwIgR0AVm6dRF7tC4Td2GzTUrtZFbbE3nuvOadLELeNEpXAKAAoBAAoLCgkKAwECAxAEGAUKIgog8crPDCZCfI46g/Qwl9tMVviJU8AJjnNsyk+fwFOPUIM=`

	block, err := ConvertToBlock(data)
	if err != nil {
		t.Fatal(err)
	}

	blockJson, err := ConvertBlockToJson(block)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(blockJson)
}