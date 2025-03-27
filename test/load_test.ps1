# Colors for output
$Green = [System.ConsoleColor]::Green
$Red = [System.ConsoleColor]::Red

Write-Host "Starting load test..." -ForegroundColor $Green

# Test 1: Basic request rate test
Write-Host "`nTest 1: Basic request rate test" -ForegroundColor Green
$messageUpdate = @{
    update_id = 123456789
    message = @{
        message_id = 123
        from = @{
            id = 123456789
            first_name = "Test User"
            username = "testuser"
            is_bot = $false
        }
        date = 1710000000
        chat = @{
            id = 123456789
            type = "private"
            first_name = "Test User"
            username = "testuser"
        }
        text = "Test message"
    }
} | ConvertTo-Json -Depth 10

bombardier -c 10 -n 1000 -m POST -b $messageUpdate http://localhost:8082/webhook

# Test 2: High concurrency test
Write-Host "`nTest 2: High concurrency test" -ForegroundColor Green
$callbackQueryUpdate = @{
    update_id = 123456790
    callback_query = @{
        id = "123456789"
        from = @{
            id = 123456789
            first_name = "Test User"
            username = "testuser"
            is_bot = $false
        }
        message = @{
            message_id = 124
            date = 1710000000
            chat = @{
                id = 123456789
                type = "private"
                first_name = "Test User"
                username = "testuser"
            }
        }
        chat_instance = "123456789"
        data = "test_callback"
    }
} | ConvertTo-Json -Depth 10

bombardier -c 50 -n 5000 -m POST -b $callbackQueryUpdate http://localhost:8082/webhook

# Test 3: Mixed update types test
Write-Host "`nTest 3: Mixed update types test" -ForegroundColor Green
$chatMemberUpdate = @{
    update_id = 123456791
    chat_member = @{
        chat = @{
            id = 123456789
            type = "group"
            title = "Test Group"
        }
        from = @{
            id = 123456789
            first_name = "Test User"
            username = "testuser"
            is_bot = $false
        }
        date = 1710000000
        old_chat_member = @{
            user = @{
                id = 123456789
                first_name = "Test User"
                username = "testuser"
                is_bot = $false
            }
            status = "member"
        }
        new_chat_member = @{
            user = @{
                id = 123456789
                first_name = "Test User"
                username = "testuser"
                is_bot = $false
            }
            status = "administrator"
            can_manage_chat = $true
            can_delete_messages = $true
            can_manage_video_chats = $true
            can_restrict_members = $true
            can_promote_members = $true
            can_change_info = $true
            can_invite_users = $true
        }
    }
} | ConvertTo-Json -Depth 10

# Start multiple background jobs for mixed update types
1..100 | ForEach-Object {
    Start-Job -ScriptBlock {
        param($messageUpdate, $callbackQueryUpdate, $chatMemberUpdate)
        $updates = @($messageUpdate, $callbackQueryUpdate, $chatMemberUpdate)
        $randomUpdate = $updates | Get-Random
        bombardier -c 1 -n 1 -m POST -b $randomUpdate http://localhost:8082/webhook
    } -ArgumentList $messageUpdate, $callbackQueryUpdate, $chatMemberUpdate
}

# Test 4: Invalid request test
Write-Host "`nTest 4: Invalid request test" -ForegroundColor Green
$invalidUpdate = @{
    invalid_field = "test"
} | ConvertTo-Json -Depth 10

bombardier -c 10 -n 100 -m POST -b $invalidUpdate http://localhost:8082/webhook

# Wait for all background jobs to complete
Get-Job | Wait-Job

# Check metrics endpoint
Write-Host "`nChecking metrics endpoint..." -ForegroundColor Green
$metrics = Invoke-RestMethod -Uri "http://localhost:8082/metrics"
$metrics | Select-String "api_requests_total|api_request_duration_seconds|telegram_updates_total|telegram_update_processing_seconds"

Write-Host "`nLoad test completed. Check Grafana dashboard at http://localhost:3000 for results." -ForegroundColor Green 