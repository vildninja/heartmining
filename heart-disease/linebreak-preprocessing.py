import re

f_in = open("cleveland-utf8.data", "r")
f_out = open("cleveland-linebreak-fixed.data", "w")

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
