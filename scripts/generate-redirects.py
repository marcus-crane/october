import requests

url = "https://api.github.com/repos/marcus-crane/october/releases/latest"
r = requests.get(url)
data = r.json()
version = data["tag_name"]

mac_redirect = f"/download/mac/latest https://github.com/marcus-crane/october/releases/download/{version}/october-darwin-universal-{version}.dmg"
win_redirect = f"/download/win/latest https://github.com/marcus-crane/october/releases/download/{version}/october-windows-amd64-{version}.dmg"

with open("site/_redirects", "w") as file:
    file.write(mac_redirect + "\n")
    file.write(win_redirect + "\n")
