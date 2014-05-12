import re
import os

# Produce joined csv header
f_in = open("../parser/HEALTH_STAT_LABELS_COUNTRY.csv", "r")
ccodes = []
for line in f_in:
	ccodes.append(line.rstrip("\n"))


def joinAttribute(headers, countries, misscount, f_in):
	cseen = dict()
	for ccode in ccodes:
		cseen[ccode] = False

	for line in f_in:
		line = line.rstrip("\n")
		attrs = line.split(',')
		name = attrs[0]
		country = countries[name]

		cseen[name] = True

		for i,attr in enumerate(attrs[1:]):
			year = headers[i+1]
			country[year].append(attr)

	for ccode in cseen.keys():
		if(not cseen[ccode]):
			misscount[ccode] = misscount[ccode] + 1
			country = countries[ccode]
			for year in country.keys():
				country[year].append("")


files = []
root = "../parser/interpolated/"
for f in os.listdir(root):
	path = root + f
	if(re.search("(?:Health_stats|Non_medical)\.xml\-([A-Z0-9]+)\-([A-Z0-9]+)\.csv", path) != None):
		if(path not in files):
			files.append(path)

# Read headers
path = files[0]
f_in = open(path, "r")
headers = f_in.next().rstrip("\n").split(',')
f_in.close()

# Prepare country dictionary to store data
countries = dict()
misscount = dict()
for ccode in ccodes:
	country = dict()
	for year in range(2000,2012):
		country[str(year)] = []
	countries[ccode] = country
	misscount[ccode] = 0

# Read and join data
for path in files:
	f_in = open(path, "r")
	f_in.next() # Skip header
	joinAttribute(headers, countries, misscount, f_in)
f_in.close()

# Remove countries with too high miss count
threshold = 120
for country in misscount.keys():
	miss = misscount[country]
	if(miss > threshold):
		print("'{0}' is missing too many features ({1}/547)".format(country, miss))
		countries.pop(country, None)

# Preview
#for ccode in countries.keys():
#	print("{0}: {1}".format(ccode, countries[ccode]))

# Produce joined csv header
f_in = open("../parser/HEALTH_STAT_LABELS.csv", "r")
labelmap = dict()
for line in f_in:
	match_c = re.search("([A-Z]+),\"?(.+)\"?", line)
	labelmap[match_c.group(1)] = match_c.group(2)
f_in.close()

f_in = open("../parser/HEALTH_STAT_LABELS_VAR.csv", "r")
for line in f_in:
	match_c = re.search("\"([A-Z0-9]+)\",\"(.+)\"", line)
	labelmap[match_c.group(1)] = match_c.group(2)
f_in.close()

f_in = open("../parser/NON_MEDICAL_LABELS.csv", "r")
for line in f_in:
	match_c = re.search("\"([A-Z0-9]+)\",\"(.+)\"", line)
	labelmap[match_c.group(1)] = match_c.group(2)
f_in.close()

f_in = open("../parser/NON_MEDICAL_LABELS_VAR.csv", "r")
varmap = dict()
for line in f_in:
	match_c = re.search("\"([A-Z0-9]+)\",\"(.+)\"", line)
	labelmap[match_c.group(1)] = match_c.group(2)
f_in.close()


headerstr = "Country;Year;"
for fl in files:
	match_c = re.search("(?:Health_stats|Non_medical)\.xml\-([A-Z0-9]+)\-([A-Z0-9]+)\.csv", fl)
	varlabel = match_c.group(1)
	label = match_c.group(2)
	detailed = "{0} [{1}]".format(labelmap[varlabel], labelmap[label])
	headerstr += "{0};".format(detailed)
	#print detailed
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
f_out = open("Health_stats_and_non_medical.xml-interpolated-merged.csv", "w")
f_out.write(headerstr)
f_out.write(csv)