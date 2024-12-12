import requests

def fetch_data(url, session_cookie):
    headers = {
        'Cookie': f'session={session_cookie}'
    }
    response = requests.get(url, headers=headers)
    response.raise_for_status()  # Raise an error for bad responses
    return response.text.strip().split('\n')

def is_safe_report(report):
    if all(report[i] < report[i + 1] and 1 <= report[i + 1] - report[i] <= 3 for i in range(len(report) - 1)):
        return True
    if all(report[i] > report[i + 1] and 1 <= report[i] - report[i + 1] <= 3 for i in range(len(report) - 1)):
        return True
    return False

def count_safe_reports(data):
    safe_count = 0
    
    for line in data:
        report = list(map(int, line.split()))
        if is_safe_report(report):
            safe_count += 1
    
    return safe_count

# URL for the input data
url = "https://adventofcode.com/2024/day/2/input"

# Your session cookie
session_cookie = "session_token"

# Fetch and process the data
data = fetch_data(url, session_cookie)
print("Number of safe reports:", count_safe_reports(data))
