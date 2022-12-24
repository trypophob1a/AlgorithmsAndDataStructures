import sys
import re
import os

args = sys.argv[1:3]
file = open(sys.argv[1], "r")
MIN_DEFAULT_PERCENT = 50

try:
    args[1]
except IndexError:
     args.append(MIN_DEFAULT_PERCENT)

try:
    min_percent = int(args[1])
except ValueError:
    print(f"\033[1m\033[91m ERROR: min percent: <<{args[1]}>> not is number:\n")
    sys.exit(1)

lines = file.readlines()
if not lines:
   print(f"\033[92m test there are no files for testing")
   command = os.popen('echo "NOT_FILE=1" >> $GITHUB_ENV')
   command.read()
   command.close()
   sys.exit(0)

file.close
compiler = re.compile(r"\d+\.\d+")
percent = float(compiler.search(lines[-1]).group())

if  percent < float(min_percent):
   print(f"\033[1m\033[91m ERROR: insufficient coverage {percent}% needs {min_percent}%\n")
   sys.exit(1)
else:
    print(f"\033[92m coverage is: {percent}%")