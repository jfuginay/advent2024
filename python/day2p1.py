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
session_cookie = "53616c7465645f5f92ad32a207f568b4e1a47eae1df9d9372fd7472e7e193069554dd925f96f21c1f801ad27cf1ae8006436eea874479919152212904690e1b5"

# Fetch and process the data
data = fetch_data(url, session_cookie)
print("Number of safe reports:", count_safe_reports(data))
