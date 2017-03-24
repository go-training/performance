with open("python", "a") as text_file:
  for i in range(1000000):
    text_file.write(str(i))
