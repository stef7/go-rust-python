import json

def get_users():
  file = open("../users.json")
  string = file.read()
  file.close()
  return json.loads(string)

def main():
  return print(get_users())
