# Quadruped Dog
Total Hours (so far): 4 hours
## july 26th, 2025
i was inspired to make this project from this article i saw: 
https://www.cnx-software.com/2024/08/08/mechdog-ai-robot-dog-features-esp32-s3-controller-supports-scratch-python-and-arduino-programming/ 

i've also always thought both boston dynamic's robot dog and MIT's mini cheetah were very cool, and i now am confident i finally have the knowledge to create a similar quadruped robot dog thanks to both designing my own drone and hackclub's undercity hackathon

this dog will be much smaller than boston dynamic's dog, currently estimating somewhere around 250mm * 150mm * 150mm, and weigh maybe around 500g

## goals/features:
- 8DOF
- traverses rough terrain (dirt)
- 30 minute battery life
- AI facial detection
- self balancing
- bluetooth controller or app driving
- autonomous pathing
- line following
- touch, temperature, lighting sensing
- voice control
- cliff detection
- obstacle detection
- USB-C self charging

more complicated features id like to implement if possible
- LIDAR terrain mapping
- walk up small stairs

tasks it should be able to accomplish:
- walk with me to the local path while i control with an xbox controller (~30 mins)
- deliver a small object to a specific person with facial detection
- voice command it to line follow to a charger and charge
- follow an autonomous path to another room in my apartment from mine
- map the terrain and send it back to my computer
- walk to my bed and wake me up with an alarm, and give me the weather
- not fall off a table

# Parts

rough parts i will need
- microcontroller
- (opt): coprocessor for LIDAR and AI
- accelerometer, gyroscope
- camera
- USB-C PD
- ultrasonic sensor
- microphone and speaker
- touch sensor
- light sensor
- battery
- LIDAR
- servo
- servo driver
- voltage regulator
- OLED screen
- (opt): leg encoders

### microcontroller
ill be using a variant of the esp32, as its a very well established microcontroller and best suited for a mini dog like this one

based on the variants, theres a few choices:

- ESP32-S3-WROOM-1-N16R8
- ESP32-WROVER-IE-N16R8
- ESP32-S3-WROOM-2-N32R8V

theres also ESP32-S3-WROOM-2-N32R16V, which is just more psram and therefore better, but it's a new product and unavailable right now so i wont design this project around it

based on the variants, i'll be going with the wroom-2-N32R8V
it has 32mb flash, 8mb PSRAM, BLE 5.0, UART, camera support, and while it is discontinued, the benefits of it outweigh this con

while looking for variants i also found this cool site to find parts: https://octopart.com/

+4 hours
### servos

deciding on coreless and digital for the faster movement and easier programming

out of all options, the Hiwonder HPS-0618SG seems to be only cheap and reasonable option for the legs (58g)

for the joints, the servo choice is less important, and so i'll be getting a much cheaper servo such as https://www.aliexpress.com/item/1005004874510884.html or https://www.aliexpress.com/item/1005008648902414.html (260g)

### battery + power

going with LiPo over Li ion because its simpler and more reliable

planning for 3S (11.1V) 3000mAh LiPo battery, which weighs around 182g

power (60W USB3.0 type-C PD):
- receptacle: UJ31-CH-3-SMT-TR
- pd, battery charge, buck control: BQ25798
- battery protection: BQ77915

### other components

- camera: OV2640 
- mic: INMP441 
- time of flight sensor: VL53L0X 
- imu: gy-87 breakout board
- touch: TTP223 
- temp: DS18B20 
- screen: SSD1309

from parts searching, deciding against these features/parts:
- encoders, LIDAR, light sensor, coprocessor

PWM driver: PCA9685 
GPIO expander: MCP23017 
servo buck converter: MP1584
3.3V buck converter: AMS1117-3.3
5V buck converter (might not be needed): LM2940
motor temperature sensor: DS18B20 

audio:
- I2S DAC: PCM5102A 
- amp: MAX98357A 
- 8Ω 1W speaker: K 20 - 8 OHM

full components parts list:

| Component                                  | Quantity | Notes                                                |
|--------------------------------------------|----------|------------------------------------------------------|
| ESP32-S3-WROOM-2-N32R8V                    | 1        | 32MB Flash, 8MB PSRAM, BLE 5.0, UART, camera support |
| 3S 11.1V 3000mAh LiPo Battery              | 1        | ~182g, main power source                             |
| USB-C Receptacle (UJ31-CH-3-SMT-TR)        | 1        | USB 3.0 PD input                                     |
| Battery Charger (BQ25798)                  | 1        | USB-C PD to LiPo charge controller                   |
| Battery Protection (BQ77915)               | 1        | 3S battery protection                                |
| Buck Converter (MP1584)                    | 1–2      | For stepping down to servo voltage (6–8.4V)          |
| 3.3V Regulator (AMS1117-3.3)               | 1        | For ESP32 and logic                                  |
| 5V Regulator (LM2940)                      | 1        | Optional, if any 5V                                  |
| Hiwonder HPS-0618SG Servos                 | 4        | Coreless digital, used for legs                      |
| Generic Digital Servos                     | 4        | For joints, cheaper option                           |
| PCA9685 PWM Driver                         | 1        | 16-channel servo controller                          |
| OV2640 Camera Module                       | 1        | For vision, facial detection, line following         |
| INMP441 Microphone                         | 1        | I2S digital mic for voice control                    |
| VL53L0X ToF Sensor                         | 1        | For object and cliff detection                       |
| GY-87 IMU Module                           | 1        | Accelerometer, gyro, magnetometer, barometer         |
| TTP223 Touch Sensor                        | 1        | Touch detection                                      |
| DS18B20 Temperature Sensor                 | 1        | For internal or external temperature sensing         |
| SSD1309 OLED Display                       | 1        | 1.3" or larger, I2C/SPI display                      |
| PCM5102A I2S DAC                           | 1        | For high-quality audio output                        |
| MAX98357A I2S Amplifier                    | 1        | Low-power mono amplifier                             |
| 8Ω 1W Speaker (K 20 - 8 OHM)               | 1        | Audio output                                         |
| MCP23017 GPIO Expander                     | 1        | I2C GPIO expansion                                   |

+4 hours