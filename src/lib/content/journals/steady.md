# Steady Journal - Total Time: 14h 30m

## 6/22/26 - 34m 59s
started the planning! it will be a linear double inverted pendulum controlled by bldc motor. i choose this because its the most common setup in youtube videos, where there is a frictionless linear rail powered by a motor, with encoders on the pivots. i also decided that if i can fit the system and a custom esc in under $180 ill do just that otherwise ill just buy an esc.

![alt text](assets/image.png)

## 6/23/26 - 1h 56m 23s
finished researching and started the cad! looked into types of linear actuators (transmission systems) and found that only two were really viable: racket and pinion system, or timing belt. the con of the timing belt was that cad-ing it would be more difficult and the belt would need to be tensioned, so i decided on racket and pinion, with the con being higher inertia of the carriage. then i looked into rail types, with aluminum extrusions with v-rails vs linear rails. linear rails are a lot more expensive ive found but it fits into budget and are more rigid and more common on designs i see, so i will go with that. after a lot of hunting and fumbling with trying to look at products on misumi, i found that hgr20 rails are the cheapest rigid rails to get. also looked into encoders and decided that my requirements for an encoder would be an absolute magnetic encoder that’s at least 12 bit. i think designing my own custom esc is rather unnecessary for this project, so i’ll stick with finding one probably. found a lot of important resources for this design: https://www.drdflo.com/pages/Guides/How-to-Build-a-3D-Printer/FFF.html https://www.youtube.com/watch?v=7eIE23vJcI4 http://willbeattie.ca/post/engineering/pendulum/

![alt text](assets/image-1.png)

## 6/26/26 - 51m 0s
figured out how to use a linear mate with the rail and guide so that part of the assembly is done. made a stand with random dimensions that ill change later. going to start cad a 400mm gear rack (dont know how to do this)

![alt text](assets/image-2.png)

## 6/26/26 - 1h 0m 1s
finished the gear rack and gear cad on the assembly. did this by learning how iso standards work to be able to replicate the dimensions in cad (except the circular radius which seemed to be different but idk??). was going to figure out the rack and pinion relation in the assembly but i actually need the motor for that. so next up is cad ing the guide mount and motor, and i guess the rod too.

![alt text](assets/image-3.png)

## 6/27/26 - 1h 25m 0s
cad ed the motor, the holder for the motor (press fit), added both to the assembly, and figured out the relations. it now has a rack and pinion relation! need to figure out the rotations thing but thats for a later time. next up is adding the rod which is going to need figuring out encoders, bearings, and the rod itself.

![alt text](assets/image-4.png)

## 6/28/26 - 1h 25m 0s
figuring out the arm attachment. its inspired from the youtube video, but mine will be all 3d prints so i can skip out and simplify a little. the idea is two brackets with a hex shaft between them and bearings. one end has a magnetic encoder to detect hex shaft rotation, and the other end will attach to a 3d print with a hex slot that will rotate. then, in the joint between arms, it will be a bearing and encoder on one end, and another 3d print with a hex slot on the other end.

![alt text](assets/image-5.png)
![alt text](assets/image-6.png)

## 6/28/26 - 2h 27m 0s
almost done cad! figured out the motion of the arms, which is with a bunch of bearings and dreams. i found out people in the real world dont actually use hex shafts but im going to stay in my fake world as long as possible so i will custom 3d print some hex shafts to work. the assembly works, and the motion all makes sense. whats left is to redo the arm part with shaft collars because those are kinda necessary, create a shorter hex shaft, create a proper stand, fix the rack and pinion relation, and then i can finally think about the electronics.

![alt text](assets/image-7.png)

## 6/28/26 - 2h 35m 0s
finished the cad! added stands to the cad that are dimensioned and fastened them, and fixed the rack and pinion relation. found a tof sensor and made a mount for that too. found a good absolute position magnetic encoder. added mounts for it on the two joints. made all the assembly relations work! next up is figuring out the hardware and wiring, and roughly the code.

![alt text](assets/image-8.png)

## 6/29/26 - 2h 16m 0s
done! i finished the BOM, then did wiring. decided on a teensy 4.1 for its very fast loop speeds and because its a cool mcu. found a motor driver and made a wiring diagram in kicad with custom symbols. i will run the encoders through i2c because the spi ones are rather massive. then figured out blender and rendered my assembly, fixing up a few things. vibecoded the minimal firmware, hopefully it works a little.

![alt text](assets/image-9.png)

