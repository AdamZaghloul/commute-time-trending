# commute-time-trending
Commute time trending is a Golang command-line tool that gets travel times from google directions api and outputs them to csv files for multiple locations going to and from a target destination.

I made this tool because my fiancee and I are looking to buy a house and exploring which neighborhoods to choose. One of the criteria is how far it is from our work. Using the Google Maps web app will give you estimated travel times at given depart/arrive-at times, but it's always a range such as 20-40 minutes. When suggesting a departure time it assumes the worst-case scenario but does not give the average travel time. This is useful for one-off travel but not for planning the average commute.

This program treats our workplace as the target destination, and gets travel times from several locations in different neighborhoods we're looking at. The tool is intended to be run as a cron job to get multiple data points that can be averaged. I run it every 30 minutes Monday to Friday to get commute times.

## 📖 Usage

Note this assumes:
1. You have golang
2. You have a .env file
    * MAPS_API_KEY - API Key for Google Maps/Directions API.
    * TARGET - The target destimation to compare all locations to. This needs to be readable by Google Maps.
    * LOCATIONS_FILE_PATH - The absolute path to the "locations.txt" file or similar with locations readable by Google Maps. One location per line.
    * TO_FILE_PATH - The absolute path to the output CSV for travel times from the locations to the target.
    * FROM_FILE_PATH - The absolute path to the output CSV for travel times from the target to the locations.
    * LOG_FILE_PATH - The absolute path to the output log.
3. You have a "locations.txt" file or similar with locations readable by Google Maps. One location per line.
4. You have two .csv files (one for times to the target and one for times from the target) with the headings you want to indicate the time and the various locations described in the locations.txt files.

### Clone the repo

```bash
git clone https://github.com/AdamZaghloul/commute-time-trending
cd commute-time-trending
```

### Run the project locally

```bash
go run .
```
You can check for success in the log.log file and if successful, see the two csv files for output.

### Submit a pull request

If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.
