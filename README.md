# CHAMP
(CSV Header Automated Merge Program)

CHAMP takes a header CSV file and merges it with a file or a folder of your choosing in a safe way. Specify the output folder and see the logs of whether the merge was successful or not in the log file. It will never overwrite files that have already been written to.

How to specify:

**-help** or **-h** This is to show you what the options for the program are

**-headerFile** string - **Default: "header.csv"** This is the file that contains the headers for the CSV file(s)

**-inputPath** string - **Default: "in"** This is the file or directory that contains the CSV file(s) you wish to merge

**-outputFolder** string - **Default: "out"** This is the directory that will contain the output files with the header

**-startingRow** int - **Default: None** This determines the file the merge starts at

**-endingRow** int - **Default: None** This determines the file the merge ends at

**-trackingFileName** string - **Default: "log.txt"** This is the file that becomes your log file for the merge
