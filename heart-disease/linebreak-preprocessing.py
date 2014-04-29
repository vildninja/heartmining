import re

# Open target output file
f_out = open("cleveland-formatted.data", "w")

# Read overall heart-disease header and add to first line
f_in = open("heart-disease-labels-fixed.data", "r")
f_out.write(f_in.next() + "\n")

# Read, parse and format data
f_in = open("cleveland-utf8.data", "r")

row_count = 0
attr_count = 0
record = ""
records = []
for line in f_in:
	vals = line.split(' ')
	for val in vals:
		match_c = re.search("name", val)
		if (match_c != None):
			record = record + "\n"
			if(attr_count == 75):
				records.append(record)
				f_out.write(record)

			record = ""
			attr_count = 0
			row_count = row_count + 1
		else:
			match_c = re.search("(-?[0-9]+)", val)
			if(match_c != None):
				record = record + match_c.group(0) + ";"
				attr_count = attr_count + 1

f_in.close()
f_out.close()

print("Row count: {0}".format(row_count))
print("Good row count: {0}".format(len(records)))
