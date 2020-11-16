Param($gitPassword, $teamcityUserId, $teamcityPwd, $target)
& docker run `
--rm `
-v "$($PSScriptRoot):C:\Source" `
-e TEAMCITY_VERSION="$($ENV:TEAMCITY_VERSION)" `
-e TEAMCITY_SERVER="$($ENV:TEAMCITY_SERVER)" `
-e TeamcityUserId="$teamcityUserId" `
-e TeamcityPwd="$teamcityPwd" `
-e Git_Password="$gitPassword" `
-e Target="$target" `
ac42dockerregistry.lmb.liebherr.i:5000/gobuild:1.15.2.1 | Out-Default
exit $LASTEXITCODE