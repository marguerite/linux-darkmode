linux-darkmode

------

Theory:

1. get your public IP address
2. use GeoLite2-City to get your longitude and latitude
3. calculate your sun rise/set time with formula here: https://en.wikipedia.org/wiki/Sunrise_equation
4. for plasma5, "lookandfeeltool -a org.kde.breeze.desktop" at sun rise and "lookandfeeltool -a org.kde.breezedark.desktop" at sun set.

That's it! All will be implemented in golang like my linux-bing-wallpaper.

Currently I only know how to switch theme on plasma5. Any help with other desktop environment are welcome!(please mention me in issues).
