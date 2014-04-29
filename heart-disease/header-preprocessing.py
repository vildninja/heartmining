import re

f_in = open("heart-disease-labels.data", "r")
f_out = open("heart-disease-labels-fixed.data", "w")

freq = dict()
first = True
for line in f_in:
	out = ""
	if(not first):
		out = ";"
	else:
		first = False
	match_c = re.search("^\s+([0-9]+)\s[a-z0-9:]+\s(.+)$", line)
	if (match_c != None):
		attrid = match_c.group(1)
		label = match_c.group(2)

		if(label in freq.keys()):
			count = freq[label]
			count = count + 1
			freq[label] = count
			label = label + str(count)
		else:
			freq[label] = 0

		out += label
		f_out.write(out)
		print("{0} : {1}".format(attrid, label))

f_in.close()
f_out.close()
