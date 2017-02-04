package script

// https://forum.justgetflux.com/topic/1928/how-to-create-keyboard-shortcuts

// Flux is the AppleScript template to execute menu option changes.
var Flux = `
property mainItem : "%s"
-- set to "Preferences...", "Color Effects", "Disable", etc.,
-- make sure to use quote marks.

property subItem : "%s"
-- set to submenu item, if there is one. Use "for this app" with Disable,
-- to toggle disable for the current application.

if mainItem is "Disable" and subItem is "for this app" then set subItem to 3

tell application "System Events"
    tell application process "Flux"
        tell menu bar 1
            tell menu bar item 1
                try
                    with timeout of 0.1 seconds
                        perform action "AXPress"
                    end timeout
                end try
            end tell
        end tell
    end tell
end tell

do shell script "killall 'System Events'"

tell application "System Events"
    tell application process "Flux"
        tell menu bar 1
            tell menu bar item 1
                tell menu 1
                    tell menu item mainItem
                        perform action "AXPress"
                    end tell
                    if menu 1 of menu item mainItem exists then
                        tell menu 1 of menu item mainItem
                            tell menu item subItem
                                perform action "AXPress"
                            end tell
                        end tell
                    end if
                end tell
            end tell
        end tell
        if mainItem is "Preferences..." then set frontmost to true
    end tell
end tell
`
