package reaction

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// DebugGesturer prints a human-readable summary of a gesturer's state.
func DebugGesturer(g *gesturer) string {
	if g == nil {
		return "<nil gesturer>"
	}

	var b strings.Builder
	colorReset := "\033[0m"
	colorCyan := "\033[36m"
	colorYellow := "\033[33m"
	colorGreen := "\033[32m"
	colorRed := "\033[31m"

	b.WriteString(fmt.Sprintf("%sGesturer Debug Info%s\n", colorCyan, colorReset))
	b.WriteString(fmt.Sprintf("%sMouse Position:%s (%d, %d)\n", colorYellow, colorReset, g.MouseX, g.MouseY))
	if g.Press != nil {
		b.WriteString(fmt.Sprintf("%sPress:%s Start(%d,%d) Current(%d,%d) Button(%d) Age(%s)\n",
			colorGreen, colorReset, g.Press.StartX, g.Press.StartY, g.Press.X, g.Press.Y, g.Press.Button, time.Since(g.Press.TimeStart).String()))
	} else {
		b.WriteString(fmt.Sprintf("%sPress:%s <none>\n", colorRed, colorReset))
	}
	b.WriteString(fmt.Sprintf("%sRegistered Events:%s\n", colorYellow, colorReset))
	if len(g.events) == 0 {
		b.WriteString("  <none>\n")
	} else {
		for typ, reactions := range g.events {
			// Sort reactions by depth, same as Register
			sorted := make([]Reaction, len(reactions))
			copy(sorted, reactions)
			sort.SliceStable(sorted, func(i, j int) bool {
				depthA := sorted[i].Depth()
				depthB := sorted[j].Depth()
				minLen := len(depthA)
				if len(depthB) < minLen {
					minLen = len(depthB)
				}
				for d := 0; d < minLen; d++ {
					if depthA[d] < depthB[d] {
						return true
					} else if depthA[d] > depthB[d] {
						return false
					}
				}
				return len(depthA) < len(depthB)
			})
			b.WriteString(fmt.Sprintf("  %s%s%s: %d reactions\n", colorCyan, typ, colorReset, len(sorted)))
			// Group reactions by Resource address
			groups := make(map[string][]Reaction)
			names := make(map[string]string)
			for _, r := range sorted {
				res := r.Resource()
				addr := "<nil>"
				name := "<none>"
				if res != nil {
					addr = fmt.Sprintf("%p", res)
					name = res.DebugName()
				}
				groups[addr] = append(groups[addr], r)
				names[addr] = name
			}
			colorMagenta := "\033[35m"
			for addr, group := range groups {
				name := names[addr]
				b.WriteString(fmt.Sprintf("    %sResource:%s %s%s%s %s[%s]%s (%d reactions)\n", colorYellow, colorReset, colorGreen, name, colorReset, colorMagenta, addr, colorReset, len(group)))
				colorGray := "\033[90m"
				for i, r := range group {
					if r.IsEnabled() {
						b.WriteString(fmt.Sprintf("      %s%d%s: enabled=%v depth=%v type=%T\n", colorGreen, i, colorReset, r.IsEnabled(), r.Depth(), r))
					} else {
						b.WriteString(fmt.Sprintf("%s      %d: enabled=%v depth=%v type=%T%s\n", colorGray, i, r.IsEnabled(), r.Depth(), r, colorReset))
					}
				}
			}
		}
	}

	// --- Additional section: All reactions per resource ---
	b.WriteString("\n" + colorCyan + "All Reactions Per Resource" + colorReset + "\n")
	// Collect all reactions by resource address
	allGroups := make(map[string][]Reaction)
	allNames := make(map[string]string)
	for _, reactions := range g.events {
		for _, r := range reactions {
			res := r.Resource()
			addr := "<nil>"
			name := "<none>"
			if res != nil {
				addr = fmt.Sprintf("%p", res)
				name = res.DebugName()
			}
			allGroups[addr] = append(allGroups[addr], r)
			allNames[addr] = name
		}
	}
	colorMagenta := "\033[35m"
	for addr, group := range allGroups {
		name := allNames[addr]
		b.WriteString(fmt.Sprintf("  %sResource:%s %s%s%s %s[%s]%s (%d reactions)\n", colorYellow, colorReset, colorGreen, name, colorReset, colorMagenta, addr, colorReset, len(group)))
		colorGray := "\033[90m"
		for i, r := range group {
			if r.IsEnabled() {
				b.WriteString(fmt.Sprintf("    %s%d%s: type=%s enabled=%v depth=%v type=%T\n", colorGreen, i, colorReset, r.ReactionType(), r.IsEnabled(), r.Depth(), r))
			} else {
				b.WriteString(fmt.Sprintf("%s    %d: type=%s enabled=%v depth=%v type=%T%s\n", colorGray, i, r.ReactionType(), r.IsEnabled(), r.Depth(), r, colorReset))
			}
		}
	}
	if g == nil {
		return "<nil gesturer>"
	}

	// var b strings.Builder
	// colorReset := "\033[0m"
	// colorCyan := "\033[36m"
	// colorYellow := "\033[33m"
	// colorGreen := "\033[32m"
	// colorRed := "\033[31m"

	// b.WriteString(fmt.Sprintf("%sGesturer Debug Info%s\n", colorCyan, colorReset))
	// b.WriteString(fmt.Sprintf("%sMouse Position:%s (%d, %d)\n", colorYellow, colorReset, g.MouseX, g.MouseY))
	// if g.Press != nil {
	// 	b.WriteString(fmt.Sprintf("%sPress:%s Start(%d,%d) Current(%d,%d) Button(%d) Age(%s)\n",
	// 		colorGreen, colorReset, g.Press.StartX, g.Press.StartY, g.Press.X, g.Press.Y, g.Press.Button, time.Since(g.Press.TimeStart).String()))
	// } else {
	// 	b.WriteString(fmt.Sprintf("%sPress:%s <none>\n", colorRed, colorReset))
	// }
	// b.WriteString(fmt.Sprintf("%sRegistered Events:%s\n", colorYellow, colorReset))
	// if len(g.events) == 0 {
	// 	b.WriteString("  <none>\n")
	// } else {
	// 	for typ, reactions := range g.events {
	// 		b.WriteString(fmt.Sprintf("  %s%s%s: %d reactions\n", colorCyan, typ, colorReset, len(reactions)))
	// 		// Group reactions by Resource address
	// 		groups := make(map[string][]Reaction)
	// 		names := make(map[string]string)
	// 		for _, r := range reactions {
	// 			res := r.Resource()
	// 			addr := "<nil>"
	// 			name := "<none>"
	// 			if res != nil {
	// 				addr = fmt.Sprintf("%p", res)
	// 				name = res.DebugName()
	// 			}
	// 			groups[addr] = append(groups[addr], r)
	// 			names[addr] = name
	// 		}
	// 		colorMagenta := "\033[35m"
	// 		for addr, group := range groups {
	// 			name := names[addr]
	// 			b.WriteString(fmt.Sprintf("    %sResource:%s %s%s%s %s[%s]%s (%d reactions)\n", colorYellow, colorReset, colorGreen, name, colorReset, colorMagenta, addr, colorReset, len(group)))
	// 			for i, r := range group {
	// 				b.WriteString(fmt.Sprintf("      %s%d%s: enabled=%v depth=%v type=%T\n", colorGreen, i, colorReset, r.IsEnabled(), r.Depth(), r))
	// 			}
	// 		}
	// 	}
	// }
	return b.String()
}
