import re
import os

def joinAttribute(headers, countries, f_in):
	for line in f_in:
		line = line.rstrip("\n")
		attrs = line.split(',')
		name = attrs[0]

		country = dict()
		if(name in countries.keys()):
			country = countries[name]
		else:
			countries[name] = country

		for i,attr in enumerate(attrs[1:]):
			year = headers[i+1]
			if(year not in country.keys()):
				country[year] = []
			country[year].append(attr)


files = []
root = "../parser/"
for f in os.listdir(root):
	path = root + f
	if(re.search("(Health_stats\.xml\-[A-Z0-9]+\-[A-Z0-9]+\.csv)", path) != None):
		if(path not in files):
			files.append(path)

# Read headers
path = files[0]
f_in = open(path, "r")
headers = f_in.next().rstrip("\n").split(',')
f_in.close()

# Read and join data
countries = dict()
for path in files:
	f_in = open(path, "r")
	f_in.next() # Skip header
	joinAttribute(headers, countries, f_in)
f_in.close()

# Preview
#for ccode in countries.keys():
#	print("{0}: {1}".format(ccode, countries[ccode]))

# Produce joined csv header
f_in = open("../parser/HEALTH_STAT_LABELS.csv", "r")
labelmap = dict()
for line in f_in:
	match_c = re.search("([A-Z]+),\"?(.+)\"?", line)
	labelmap[match_c.group(1)] = match_c.group(2)

f_in2 = open("../parser/HEALTH_STAT_LABELS_VAR.csv", "r")
varmap = dict()
for line in f_in2:
	match_c = re.search("\"([A-Z0-9]+)\",\"(.+)\"", line)
	varmap[match_c.group(1)] = match_c.group(2)


headerstr = "\"Country\";\"Year\";"
for fl in files:
	match_c = re.search("Health_stats\.xml\-([A-Z0-9]+)\-([A-Z0-9]+)\.csv", fl)
	varlabel = match_c.group(1)
	label = match_c.group(2)
	detailed = "{0} [{1}]".format(varmap[varlabel], labelmap[label])
	headerstr += "\"{0}\";".format(detailed)
headerstr = headerstr.rstrip(";")
headerstr += "\n"

# Produce joined csv
csv = ""
for ccode in countries.keys():
	country = countries[ccode]
	for year in country.keys():
		csv += "{0};{1};".format(ccode, year)
		for val in country[year]:
			csv += "{0};".format(val)
		csv = csv.rstrip(";")
		csv += "\n"

# Output to csv
f_out = open("Health_stats.xml-merged.csv", "w")
f_out.write(headerstr)
f_out.write(csv)