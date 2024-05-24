import requests

url = "https://api.github.com/repos/marcus-crane/october/releases/latest"
r = requests.get(url)
data = r.json()
version = data["tag_name"]
normalised_version = version.replace("v", "")

mac_redirect = f"/download/mac/latest https://github.com/marcus-crane/october/releases/download/{version}/october_{normalised_version}_darwin_universal.zip"
win_redirect = f"/download/win/latest https://github.com/marcus-crane/october/releases/download/{version}/october_{normalised_version}_windows_amd64.zip"
win_portable_redirect = f"/download/win-portable/latest https://github.com/marcus-crane/october/releases/download/{version}/october_{normalised_version}_windows-portable_amd64.zip"
linux_redirect = f"/download/linux/latest https://github.com/marcus-crane/october/releases/download/{version}/october_{normalised_version}_linux_amd64.zip"

with open(".vitepress/dist/_redirects", "w") as file:
    file.write(mac_redirect + "\n")
    file.write(win_redirect + "\n")
    file.write(win_portable_redirect + "\n")
    file.write(linux_redirect + "\n")
