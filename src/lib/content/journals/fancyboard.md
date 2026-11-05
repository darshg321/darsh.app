---
title: "FancyBoard"
author: "Darsh Gupta"
description: "A custom, low-profile, hot-swappable mechanical keyboard."
created_at: "2025-07-28"
---

## time spent: 36 hours
## july 28th

speedrunning a keyboard before the 31st, just because

i was inspired for making this keyboard by the KNOB keyboard, which ive seen and always thought looked like the perfect keyboard. unfortunately, its $439, bad, so i'll try to make something similar for $150

KNOB:
![KNOB](assets/image1.png)

planned features:
- nice mechanical switches
- keycaps, i wont be using low profile because i personally dont like it
- hotswappable
- 2 dials
- oled screen
- both usb c and bluetooth connectivity
- usb c charging
- 3000mAh battery for big battery life

ill be using the nice!nano v2 because it consumes very little power, has battery charging built in, and has bluetooth HID + usb c HID already

parts list:

| Component                 | Part                               | Notes                                             |
| ------------------------- | ---------------------------------- | ------------------------------------------------- |
| Microcontroller           | Nice!Nano V2                       | nRF52840 + integrated battery charger             |
| Battery                   | 3.7V Li-ion 3000mAh (flat, JST-PH) | Rechargeable                                      |
| USB Port                  | JST to USB-C charging cable        | Board already has onboard USB                     |
| Rotary Encoder            | EC11 w/ push button                | For volume/menu                                   |
|     Screen                | ST7789 1.14"                       | SPI, vertical     https://www.aliexpress.com/item/1005009445748262.html                                |

| Component       | Part                                        | Notes                 |
| --------------- | ------------------------------------------- | --------------------- |
| Switches        | 78x Tactile MX Switches                     |                       |
| Diodes          | 1N4148W SMD                                 |                       |
| Hotswap Sockets | 78x Kailh MX sockets                        |                       |
| Keycaps         | SA / DSA / OEM profile                      | PC or PBT preferred   |
| Encoder Knob    | Aluminum / ABS cap                          |                       |

# schematic

![2](assets/image2.png)

microcontroller symbol done!

+3 hours

![3](assets/image3.png)

done with ST7789_Module symbol

![4](assets/image4.png)
current progress of schematic

+2 hours

adding 2 MCP23017 because i need a lot more io

## july 29th

![5](assets/image5.png)

added all the components, need to complete wiring

![6](assets/image6.png)

finished most of everything, need to add some pull up resistors and capacitors and stuff

+3 hours

![7](assets/image7.png)

done with schematic, pull ups and capacitors

need to make footprints and assignments

![8](assets/image8.png)

![9](assets/image9.png)

done with both footprints

+3 hours

keycaps: https://www.aliexpress.com/item/1005004840360158.html

deciding on Kailih Choc V2 Low Profile Switches

stabalizers: https://www.aliexpress.com/item/1005001781468636.html

spdt switch: SS12D00

![10](assets/image10.png)

fixed a few symbol and footprint issues, starting pcb

switching to choc v1 because its easier
https://shop.beekeeb.com/product/kailh-choc-v1-hotswap-socket/


![11](assets/image11.png)

fixed the layout, need to route now

![12](assets/image12.png)
finshed layout fully, routing time

+5 hours

## july 30th

![13](assets/image13.png)

almost finished routing, need to route rows/columns and other components

finished routing rows and columns
![14](assets/image14.png)

apparently the footprint library had the sockets flipped, so i had to edit the footprint library and then reroute all the keys unfortunately

cooked but oh well

![15](assets/image15.png)

+4 hours

![16](assets/image16.png)

completely finished pcb, working on models

+2 hours

changed switch, cleaned up some wiring, added 3d models for everything, changed some footprints

pcb fully done !!

![](assets/schematic.png)

![](assets/pcb.png)

![](assets/viewer.png)

+2 hours

## july 31st

manually adding keys cause kicad doesnt export stretched components properly (its sooooo slow)

![](assets/image17.png)

also made the bom

+3 hours

![](assets/image18.png)

done with pcb cad

![](assets/image19.png)

top plate cad done

![](assets/image20.png)

bottom plate mostly done, need to add usb cutout still

![](assets/image21.png)

working on case

![](assets/image22.png)

cad mostly done need to add some design later

+4 hours

firmware:

based it on the bt75_v1 firmware, as its very similar

using https://nickcoutsos.github.io/keymap-editor/, i edited a zmk keymap for the keyboard

![](assets/image23.png)

after some chatgpt consulting and perusing the docs, zmk is actually really complicated and lowkey a little gibberish, so im gonna switch over to bluemicron firmware, which also supports the nicenanov2 and has bluetooth, but uses arduino c instead of the mess zmk has

learning bluemicro: copied over repo and made my project, starting on config

![](assets/image24.png)

finished code, mostly default but it should work

changed a bunch of BOM stuff to make pricing work

+2 hours

added cad designs
![](assets/image25.png)
![](assets/image26.png)

done !!

![](assets/cad.png)

added render, filleted corners, assigned appearance to everything

![](assets/image27.png)

added holes w fillet

+2 hours

![](assets/image28.png)
changing design

![](assets/image29.png)

![](assets/image30.png)
finished editing plate

+1 hour

![](assets/image31.png)
added mounting holes, m2 screws (hole is 18mm down)