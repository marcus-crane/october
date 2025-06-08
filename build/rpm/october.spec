Name:           october
Version:        0.0.0
Release:        1%{?dist}
Summary:        A desktop application that can send your Kobo highlights to Readwise

License:        MIT
URL:            https://october.utf9k.net
Source0:        october
BuildArch:      ARCH

# Dependencies for different distributions
%if 0%{?fedora} >= 35 || 0%{?rhel} >= 9
Requires:       gtk3
Requires:       webkit2gtk4.1
%else
Requires:       gtk3
Requires:       webkit2gtk3
%endif

%description
A small Wails application for syncing your highlights with Readwise across Windows, macOS and Linux.
Kobo eReaders are somewhat notorious for not being user friendly to extract highlights off of. 
Personally I only send mine to Readwise anyway so this tool does just that in as little as two clicks.

%prep
# No prep needed for pre-built binary

%build
# No build needed for pre-built binary

%install
rm -rf %{buildroot}
mkdir -p %{buildroot}/usr/local/bin
mkdir -p %{buildroot}/usr/share/applications
mkdir -p %{buildroot}/usr/share/icons/hicolor/16x16/apps
mkdir -p %{buildroot}/usr/share/icons/hicolor/32x32/apps
mkdir -p %{buildroot}/usr/share/icons/hicolor/48x48/apps
mkdir -p %{buildroot}/usr/share/icons/hicolor/64x64/apps
mkdir -p %{buildroot}/usr/share/icons/hicolor/128x128/apps
mkdir -p %{buildroot}/usr/share/icons/hicolor/256x256/apps
mkdir -p %{buildroot}/usr/share/icons/hicolor/512x512/apps
mkdir -p %{buildroot}/usr/share/metainfo

# Install binary
install -m 755 %{SOURCE0} %{buildroot}/usr/local/bin/october

# Install desktop file
cat > %{buildroot}/usr/share/applications/october.desktop << 'EOF'
[Desktop Entry]
Name=October
Exec=/usr/local/bin/october %U
Terminal=false
Type=Application
Icon=october
StartupWMClass=october
Comment=A tool for syncing Kobo highlights to Readwise
MimeType=x-scheme-handler/october;
Categories=Office;
EOF

# Install metainfo file
cat > %{buildroot}/usr/share/metainfo/net.utf9k.october.appdata.xml << 'EOF'
<?xml version="1.0" encoding="UTF-8"?>
<!-- Copyright 2018-2019 utf9k -->
<component type="desktop-application">
  <id>net.utf9k.october</id>
  <metadata_license>CC0-1.0</metadata_license>
  <project_license>MIT</project_license>
  <developer_name>Marcus Crane</developer_name>
  <update_contact>marcus@utf9k.net</update_contact>
  <url type="homepage">https://october.utf9k.net</url>
  <url type="bugtracker">https://github.com/marcus-crane/october/issues</url>
  <name>October</name>
  <summary>Send your Kobo highlights to Readwise in two clicks.</summary>
  <description>
    <p>Getting highlights off of your Kobo is quite fiddly on a technical level.</p>
    <p>October is a community-driven desktop application that makes it really simple to send them to Readwise.</p>
    <p>100% open source with support for Windows, macOS and Linux!</p>
  </description>
  <categories>
    <category>Office</category>
  </categories>
  <launchable type="desktop-id">october.desktop</launchable>
  <icon type="remote" height="512" width="512">https://raw.githubusercontent.com/marcus-crane/october/main/build/linux/october_0.0.0_ARCH/usr/share/icons/hicolor/512x512/apps/october.png</icon>
  <content_rating type="oars-1.0" />
  <project_group>utf9k</project_group>
  <requires>
    <internet>always</internet>
  </requires>
</component>
EOF

# Install default icon (will be overridden if specific size icons are available)
install -m 644 %{_topdir}/SOURCES/appicon.png %{buildroot}/usr/share/icons/hicolor/512x512/apps/october.png 2>/dev/null || echo "512x512 icon not found, will use fallback"
install -m 644 %{_topdir}/SOURCES/appicon.png %{buildroot}/usr/share/icons/hicolor/256x256/apps/october.png 2>/dev/null || echo "256x256 icon not found, will use fallback"
install -m 644 %{_topdir}/SOURCES/appicon.png %{buildroot}/usr/share/icons/hicolor/128x128/apps/october.png 2>/dev/null || echo "128x128 icon not found, will use fallback"
install -m 644 %{_topdir}/SOURCES/appicon.png %{buildroot}/usr/share/icons/hicolor/64x64/apps/october.png 2>/dev/null || echo "64x64 icon not found, will use fallback"
install -m 644 %{_topdir}/SOURCES/appicon.png %{buildroot}/usr/share/icons/hicolor/48x48/apps/october.png 2>/dev/null || echo "48x48 icon not found, will use fallback"
install -m 644 %{_topdir}/SOURCES/appicon.png %{buildroot}/usr/share/icons/hicolor/32x32/apps/october.png 2>/dev/null || echo "32x32 icon not found, will use fallback"
install -m 644 %{_topdir}/SOURCES/appicon.png %{buildroot}/usr/share/icons/hicolor/16x16/apps/october.png 2>/dev/null || echo "16x16 icon not found, will use fallback"

%post
# Update desktop database and icon cache
if [ -x /usr/bin/update-desktop-database ]; then
    /usr/bin/update-desktop-database -q /usr/share/applications
fi
if [ -x /usr/bin/gtk-update-icon-cache ]; then
    /usr/bin/gtk-update-icon-cache -q /usr/share/icons/hicolor
fi

%postun
# Update desktop database and icon cache after uninstall
if [ $1 -eq 0 ]; then
    if [ -x /usr/bin/update-desktop-database ]; then
        /usr/bin/update-desktop-database -q /usr/share/applications
    fi
    if [ -x /usr/bin/gtk-update-icon-cache ]; then
        /usr/bin/gtk-update-icon-cache -q /usr/share/icons/hicolor
    fi
fi

%files
/usr/local/bin/october
/usr/share/applications/october.desktop
/usr/share/metainfo/net.utf9k.october.appdata.xml
/usr/share/icons/hicolor/*/apps/october.png

%changelog
* Mon Jun 02 2025 Marcus Crane <marcus@utf9k.net> - 0.0.0-1
- Initial RPM package for October
- Added support for Fedora and RHEL distributions
- Includes desktop integration and icon cache updates 