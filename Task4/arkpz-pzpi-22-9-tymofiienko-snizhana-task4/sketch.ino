#include "DHTesp.h"
#include <Adafruit_MPU6050.h>
#include <WiFi.h>
#include <HTTPClient.h>
#include "time.h"

const int DHT_PIN = 15;

DHTesp dhtSensor;
Adafruit_MPU6050 mpu;
#define MPU6050_ADDR 0x68;

const char* ssid = "Wokwi-GUEST";
const char* password = "";

const char* serverUrl = "http://2816-82-144-214-50.ngrok-free.app/api/pet-and-health/v1/health-data/device";

const char* ntpServer = "pool.ntp.org";
const long gmtOffset_sec = 0;
const int daylightOffset_sec = 0;

const char* petId = "6746db1e7a2137a0f967604d";
float activity;
unsigned long lastUpdate = 0;

unsigned long sleepStartTime = 0;
unsigned long totalSleepTime = 0;
bool isSleeping = false;

void setup() {
  Serial.begin(115200);
  dhtSensor.setup(DHT_PIN, DHTesp::DHT22);

  if (!mpu.begin()) {
    Serial.println("Failed to find MPU6050 chip");
    while (1) {
      delay(10);
    }
  }

  mpu.setAccelerometerRange(MPU6050_RANGE_8_G);
  mpu.setGyroRange(MPU6050_RANGE_250_DEG);
  mpu.setFilterBandwidth(MPU6050_BAND_21_HZ);

  Serial.println("Connecting to Wi-Fi...");
  WiFi.begin(ssid, password);

  while (WiFi.status() != WL_CONNECTED) {
    delay(2000);
    Serial.println("Current Wi-Fi Status: " + String(WiFi.status()));

  }
  Serial.println("Wi-Fi connected!");

  configTime(gmtOffset_sec, daylightOffset_sec, ntpServer);
  Serial.println("Time synchronized.");
}

void sendDataToServer(float temp, float activity, float sleepHours) {
  String jsonData = "{";
  jsonData += "\"pet_id\":\"" + String(petId) + "\",";
  jsonData += "\"activity\":" + String(activity, 2) + ",";
  jsonData += "\"sleep_hours\":" + String(sleepHours, 2) + ",";
  jsonData += "\"temperature\":" + String(temp, 2);
  jsonData += "}";

  Serial.print("JSON: ");
  Serial.println(jsonData);

  if (WiFi.status() == WL_CONNECTED) {
    HTTPClient http;
    http.begin(serverUrl);

    http.addHeader("Content-Type", "application/json");
    int httpCode = http.POST(jsonData);

    Serial.print("HTTP POST Code: ");
    Serial.println(httpCode);

  if (httpCode > 0) {
    String response = http.getString();
    Serial.print("Server Response: ");
    Serial.println(response);
  } else {
    Serial.println("Error in HTTP request");
  }

  http.end();
  } else {
  Serial.println("WiFi Disconnected");
  }
}

void loop() {
  struct tm timeinfo;
  if (!getLocalTime(&timeinfo)) {
    Serial.println("Failed to obtain time");
    return;
  }

  if (timeinfo.tm_hour == 0 && timeinfo.tm_min == 0) {
    totalSleepTime = 0;
    Serial.println("Sleep time reset to 0 (new day started).");
    delay(60000);
  }

  TempAndHumidity  data = dhtSensor.getTempAndHumidity();
  Serial.println("Temp: " + String(data.temperature, 2) + "°C");

  sensors_event_t a, g, temp;
  mpu.getEvent(&a, &g, &temp);

  Serial.print("Accel X: ");
  Serial.print(a.acceleration.x);
  Serial.print(" Y: ");
  Serial.print(a.acceleration.y);
  Serial.print(" Z: ");
  Serial.println(a.acceleration.z);

  activity = sqrt(a.acceleration.x * a.acceleration.x +
                  a.acceleration.y * a.acceleration.y +
                  a.acceleration.z * a.acceleration.z);

  Serial.print("Activity: ");
  Serial.println(activity);

  if (activity < 5) {
    if (!isSleeping) {
      isSleeping = true;
      sleepStartTime = millis();
    }
  } else {
    if (isSleeping) {
      totalSleepTime += millis() - sleepStartTime;
      isSleeping = false;
    }
  }

  Serial.print("Sleep hours: ");
  Serial.println(totalSleepTime/ 3600000.0);
  Serial.println("---");

  if (millis() - lastUpdate > 10000) {
    lastUpdate = millis();

    float sleepHours = totalSleepTime / 3600000.0;

    sendDataToServer(data.temperature, activity, sleepHours);
  }

  delay(2000);
}