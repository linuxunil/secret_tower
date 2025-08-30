from src.lib import loader

def main():
    user_input = ""
    map = loader.load_map("src/dist/map.toml")

    # game loop
    while user_input not in ("quit", "exit"):
        user_input = input("Zork> ")
        if user_input == "map":
            for room in map["rooms"]:
                print(room["name"])
                print(room["description"])
                for exit in room["exits"]:
                    print(f" -> {exit['room']}")



if __name__ == "__main__":
    main()
