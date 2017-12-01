# compute the angles in radians of the hands of the clock
now = datetime.datetime.now()
hours = now.hour * math.pi / 6
minutes = now.minute * math.pi / 30
seconds = now.second * math.pi / 30

context.set_line_cap(cairo.LINE_CAP_ROUND)

# draw the hours hand
context.save()
context.set_source_rgba(0.337, 0.612, 0.117, 0.9) # green
context.set_line_width(7)
context.move_to(x, y)
context.line_to(x + math.sin(hours + minutes/12) * (radius * 0.5),
                y - math.cos(hours + minutes/12) * (radius * 0.5))
context.stroke()
context.restore()

# draw the minutes hand
context.save()
context.set_source_rgba(0.117, 0.337, 0.612, 0.9) # blue
context.set_line_width(5)
context.move_to(x, y)
context.line_to(x + math.sin(minutes + seconds/60) * (radius * 0.8),
                y - math.cos(minutes + seconds/60) * (radius * 0.8))
context.stroke()
context.restore()

# draw the seconds hand
context.save()
context.set_source_rgba(0.7, 0.7, 0.7, 0.8) # gray
context.set_line_width(3)
context.move_to(x, y)
context.line_to(x + math.sin(seconds) * (radius * 0.9),
                y - math.cos(seconds) * (radius * 0.9))
context.stroke()
context.restore()