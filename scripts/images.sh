#!/usr/bin/env bash
# Builds static/img from the archived catalogue. Re-run after changing the manifest below.
# Each line: output path (under static/img, no extension) | source (under archive/) | max width
set -euo pipefail
cd "$(dirname "$0")/.."
A=archive
OUT=static/img
R=content/repos
C=content/contributions
M=content/remote

while IFS='|' read -r out src w; do
  [[ -z "$out" || "$out" == \#* ]] && continue
  dest="$OUT/$out.webp"
  mkdir -p "$(dirname "$dest")"
  if [[ ! -f "$dest" || "$A/$src" -nt "$dest" ]]; then
    # cutouts get their empty margins trimmed so they sit evenly on the boards
    trim=(); [[ "$out" == cut/* ]] && trim=(-fuzz 3% -trim +repage)
    convert "$A/$src[0]" -auto-orient "${trim[@]}" -resize "${w}x${w}>" -strip -quality 82 -define webp:alpha-quality=90 "$dest"
    echo "built $dest"
  fi
done <<EOF
# ---------- cutouts (card covers) ----------
cut/icey|concept/img/icey.jpg|900
cut/icey-render|concept/img/icey-render.jpg|1000
cut/ultisense|concept/img/ultisense.png|700
cut/easyrp2040|concept/img/easyrp2040.png|400
cut/spotiplay|concept/img/spotiplay.png|700
cut/fume|concept/img/fume.png|900
cut/arcadetris|concept/img/arcadetris.png|900
cut/toparm|concept/img/toparm.png|600
cut/aura|concept/img/aura.png|300
cut/badge|concept/img/badge.png|520
cut/hackpad|concept/img/hackpad.jpg|600
cut/minicopter|concept/img/minicopter.png|900
cut/miniair|concept/img/miniair.png|900
cut/steady|concept/img/steady.png|900
cut/boreas|concept/img/boreas.jpg|900
cut/keychaintris|concept/img/keychaintris.jpg|900
cut/crane|concept/img/crane.png|900
cut/deadline|concept/img/deadline.png|900
cut/fancyboard|concept/img/fancyboard.png|900
cut/mushu|concept/img/mushu.png|900
# ---------- project galleries ----------
p/icey/photo|$R/Icey/images/learn__public__img__icey-photo.webp|1400
p/icey/render|$R/Icey/images/assets__RENDER.jpg|1400
p/icey/render2|$R/Icey/images/assets__RENDER_2.jpg|1400
p/icey/back|$R/Icey/images/assets__3DPCBBACK.png|1400
p/icey/pcb|$R/Icey/images/assets__PCB.png|1400
p/icey/schematic|$R/Icey/images/assets__SCHEMATIC.png|1600
p/icey/vga-plasma|$R/Icey/images/firmware__examples__icestudio__vga_demo_preview__1_plasma.png|800
p/icey/vga-life|$R/Icey/images/firmware__examples__icestudio__showcase_vga_preview__1_life_rainbow.png|800
p/icey/vga-256|$R/Icey/images/firmware__examples__icestudio__showcase_vga_preview__4_256_colours.png|800
p/ultisense/build|$R/Ultisense/images/assets__BUILD.PNG|1200
p/ultisense/pcb3d|$R/Ultisense/images/assets__PCB_3D.png|1200
p/ultisense/pcb|$R/Ultisense/images/assets__PCB.png|1200
p/ultisense/schematic|$R/Ultisense/images/assets__Schematic.png|1600
p/easyrp2040/build|$R/EasyRP2040/images/BUILD.png|1000
p/easyrp2040/photo|$R/EasyRP2040/images/readme__7779f488-6cd1-4d6c-a19a-be1e9f82466a.jpg|1400
p/easyrp2040/pcb|$R/EasyRP2040/images/EasyRP2040.png|1000
p/spotiplay/build|$R/Spotiplay/images/BUILD.png|1200
p/spotiplay/cad|$R/Spotiplay/images/CAD.PNG|1200
p/spotiplay/wiring|$R/Spotiplay/images/WIRING.png|1200
p/fume/build|$R/Fume-Extractor/images/assets__BUILD.png|1200
p/fume/poster|$R/Fume-Extractor/images/assets__POSTER.png|1200
p/fume/render|$R/Fume-Extractor/images/assets__RENDER.png|1400
p/fume/cad|$R/Fume-Extractor/images/assets__CAD.png|1200
p/fume/wiring|$R/Fume-Extractor/images/assets__WIRING.png|1200
p/arcadetris/working|$R/arcadetris/images/images__working.jpeg|1400
p/arcadetris/posture|$R/arcadetris/images/images__peakposture.jpeg|1400
p/arcadetris/posture2|$R/arcadetris/images/images__morepeakposture.jpeg|1400
p/arcadetris/start|$R/arcadetris/images/readme__5b415f23-345b-4fd1-97a2-3a29e4897af6.jpg|1400
p/arcadetris/build|$R/arcadetris/images/readme__c40e6ec5-4b0b-41e1-881a-fba4a8f5ad39.png|1400
p/toparm/build|$R/TopArm/images/BUILD.jpg|1200
p/toparm/concept|$R/TopArm/images/CONCEPT.png|1200
p/aura/build|$R/aura-generator/images/BUILD.png|1000
p/aura/concept|$R/aura-generator/images/CONCEPT.PNG|1200
p/badge/build|$R/fallout-badge/images/BUILD.png|1000
p/badge/pcb|$R/fallout-badge/images/PCB-1.png|1000
p/hackpad/built|$R/darshs-hackpad/images/built.jpg|1200
p/hackpad/render|$R/darshs-hackpad/images/readme__28f5fd32-bb48-4e42-9482-5625a47d162f.png|1200
p/hackpad/pcb|$R/darshs-hackpad/images/readme__ba89564c-b953-4425-9036-cf7bcac6190f.png|1200
p/minicopter/render|$R/Minicopter/images/assets__Minicopter_Render.png|1400
p/minicopter/cad|$R/Minicopter/images/assets__cad.png|1400
p/minicopter/pcb3d|$R/Minicopter/images/assets__3dpcb.png|1400
p/minicopter/board1|$R/Minicopter/images/assets__image29.jpg|1200
p/minicopter/board2|$R/Minicopter/images/assets__image30.jpg|1200
p/minicopter/board3|$R/Minicopter/images/assets__image31.jpg|1200
p/minicopter/board4|$R/Minicopter/images/assets__image33.jpg|1200
p/minicopter/bench|$R/Minicopter/images/assets__image34.JPG|1200
p/miniair/fc|$R/Mini-Air/images/assets__PCB_RENDER.png|1200
p/miniair/fc2|$R/Mini-Air/images/assets__PCB_RENDER2.png|1200
p/miniair/drone|$R/Mini-Air/images/assets__CAD_RENDER2.png|1400
p/miniair/poster|$R/Mini-Air/images/assets__POSTER.png|1200
p/miniair/pcb|$R/Mini-Air/images/assets__PCB.png|1200
p/miniair/schematic|$R/Mini-Air/images/assets__SCHEMATIC.png|1600
p/steady/render|$R/Steady/images/assets__RENDER.png|1400
p/steady/cad|$R/Steady/images/assets__CAD.png|1400
p/steady/planning|$R/Steady/images/assets__planning.png|1400
p/steady/wiring|$R/Steady/images/assets__WIRING.png|1400
p/boreas/poster|$R/Boreas/images/assets__POSTER.png|1200
p/boreas/render|$R/Boreas/images/assets__RENDER.jpg|1400
p/boreas/pcb-jlc|$R/Boreas/images/assets__PCB_JLC.png|1200
p/boreas/pcb|$R/Boreas/images/assets__PCB.png|1200
p/boreas/schematic|$R/Boreas/images/assets__SCHEMATIC.png|1600
p/keychaintris/poster|$R/Keychain-tris/images/assets__POSTER.png|1200
p/keychaintris/render|$R/Keychain-tris/images/assets__RENDER.jpg|1400
p/keychaintris/pcb|$R/Keychain-tris/images/assets__PCB.png|1200
p/keychaintris/schematic|$R/Keychain-tris/images/assets__SCHEMATIC.png|1600
p/crane/poster|$R/Crane-Advantage/images/assets__POSTER.png|1200
p/crane/concept|$R/Crane-Advantage/images/assets__CONCEPT.png|1200
p/crane/cad|$R/Crane-Advantage/images/assets__CAD.png|1200
p/crane/wiring|$R/Crane-Advantage/images/assets__WIRING.png|1200
p/deadline/cad|$R/Deadline/images/assets__CAD.png|1400
p/deadline/plan|$R/Deadline/images/assets__PLAN.png|1400
p/deadline/wiring|$R/Deadline/images/assets__WIRING.png|1200
p/fancyboard/render|$R/FancyBoard/images/assets__FancyBoard_Render.png|1400
p/fancyboard/cad|$R/FancyBoard/images/assets__cad.png|1400
p/fancyboard/pcb|$R/FancyBoard/images/pcb__FancyBoard.png|1400
p/bingbang/gameplay|$C/flyingf15h__bing-bang-dream-adventure/images/readme__60197b6f-13b8-43c3-ab1e-5cd978b65b20.png|1400
p/bingbang/intro|$C/flyingf15h__bing-bang-dream-adventure/images/readme__1a9cd251-4646-421f-840c-63d1d7461880.png|1400
p/bingbang/select|$C/flyingf15h__bing-bang-dream-adventure/images/readme__ba1901ba-3d48-42c2-ab56-feaaa6430755.png|1400
p/bingbang/results|$C/flyingf15h__bing-bang-dream-adventure/images/readme__6856caf3-e814-4170-a9eb-1257b985db98.png|1400
p/bingbang/leaderboard|$C/flyingf15h__bing-bang-dream-adventure/images/readme__3498af87-1977-4f82-af06-008c83f14767.png|1400
p/jollipop/build|$C/penguinissus__jollynator/images/assets__BUILD.png|900
p/jollipop/poster|$C/penguinissus__jollynator/images/assets__POSTER.png|1200
p/jollipop/wiring|$C/penguinissus__jollynator/images/assets__WIRING.png|1200
p/diarymuse/1|$C/JasonChou0105__DiaryMuse/images/readme__1947ed5b158335071bf43944724bdfd7.png|1400
p/diarymuse/2|$C/JasonChou0105__DiaryMuse/images/readme__1947eec9618bada0b2b292a418a8c4df.png|1400
p/diarymuse/3|$C/JasonChou0105__DiaryMuse/images/readme__1947ef006396bd3ba3d30844e3db97fb.png|1400
p/productite/1|$R/Productite/images/Screenshots_Video__Screenshot_1735432179.png|600
p/productite/2|$R/Productite/images/Screenshots_Video__Screenshot_1735432185.png|600
p/productite/3|$R/Productite/images/Screenshots_Video__Screenshot_1735432188.png|600
p/bench/demo|$R/bench/images/DEMO.png|1400
p/winconvert/demo|$R/Win-Convert/images/assets__demo.png|1400
p/undyne/battle|$R/Undyne-Simulator/images/assets__images__undyne_battle.gif|800
p/sprig/back|$C/hackclub__sprig/images/readme__sprig-back.jpeg|1000
p/helpconnect/post|$R/HelpConnect/images/public__createpost.png|1200
p/helpconnect/logo|$R/HelpConnect/images/public__logo.png|600
# ---------- event + org assets ----------
e/fh27-logo|$M/fh27-logo.webp|900
e/fh26-logo|$M/fh26-logo.png|900
e/cs-logo|$M/cs-logo.webp|400
e/dayoneof-wordmark|$M/dayoneof-wordmark.webp|900
e/dayoneof-reel|$M/dayoneof-reel-darsh.jpg|600
e/undercity-gates|$M/undercity-gates.png|1600
e/outpost-logo|$M/outpost-logo.png|900
e/fallout-clouds|$M/fallout-clouds.png|1600
e/rotary|$M/fh26-sponsor-Rotary.png|400
e/ap|$M/fh26-sponsor-AP.png|400
me/darsh|$M/fh26-darsh.jpg|700
e/fh24/1|$M/fh24-1.jpeg|1400
e/fh24/3|$M/fh24-3.jpeg|1400
e/fh24/4|$M/fh24-4.jpeg|1400
e/fh24/6|$M/fh24-6.jpeg|1400
e/fh26/leads|$M/fh26-0010.jpg|1400
e/fh26/stage|$M/fh26-0019.jpg|1000
e/fh26/stage2|$M/fh26-0039.jpg|1400
e/fh26/rotary|$M/fh26-0026.jpg|1000
e/fh26/hackers|$M/fh26-0060.jpg|1400
e/fh26/merch|$M/fh26-0083.jpg|1400
e/fh26/mentoring|$M/fh26-0101.jpg|1400
e/fh26/tables|$M/fh26-0123.jpg|1400
e/fh26/hall|$M/fh26-0149.jpg|1400
e/cs/proj-fh26|$M/cs-proj-fh26.webp|800
e/cs/proj-fh27|$M/cs-proj-fh27.webp|800
e/cs/proj-dayoneof|$M/cs-proj-dayoneof.webp|800
e/cs/meet1-1|$M/cs-meet1-haile_godly.webp|600
e/cs/meet1-2|$M/cs-meet1-arvin_tanish.webp|600
e/cs/meet1-3|$M/cs-meet1-louis_suman.webp|600
e/cs/meet1-4|$M/cs-meet1-natalie_arlie.webp|600
e/cs/meet1-5|$M/cs-meet1-penny_taemin.webp|600
e/cs/meet1-6|$M/cs-meet1-yusuf.webp|600
e/offscript/crowd|$M/offscript-crowd.jpg|1000
e/offscript/room|$M/offscript-room.jpg|1000
e/offscript/demo|$M/offscript-demo.jpg|1000
e/offscript/talk|$M/offscript-talk.jpg|1000
# ---------- drawn for projects with no pictures ----------
p/claudle/chat|$M/gen/claudle.png|1200
p/guessrgo/map|$M/gen/guessrgo.png|1200
p/face/frame|$M/gen/face.png|1200
# ---------- borrowed from each event's own site ----------
e/polaris-bg|$M/themes/polaris-bg.webp|1600
e/hackaday-logo|$M/themes/hackaday-logo.png|400
s/htn|$M/themes/shot-htn.png|1400
s/hackaday|$M/themes/shot-hackaday.png|1400
s/uottahack|$M/themes/shot-uottahack.png|1400
e/htn-art|$M/themes/htn-art.png|900
e/wtw-skyline|$M/themes/wtw-skyline-halftone.png|1400
e/offscript/table|$M/offscript-photos/DSC_0098.jpg|1400
e/offscript/pizza|$M/offscript-photos/DSC_0040.jpg|1400
e/offscript/chairs|$M/offscript-photos/DSC_0028.jpg|1400
e/offscript/demo2|$M/offscript-photos/DSC_0090.jpg|1400
e/offscript/whiteboard|$M/offscript-photos/DSC_0118.jpg|1400
e/offscript/plates|$M/offscript-photos/DSC_0121.jpg|1400
e/offscript/wide|$M/offscript-photos/DSC_0013.jpg|1400
e/offscript/host|$M/offscript-host-crop.png|900
e/uottahack-art|$M/themes/uottahack-art.png|900
e/fallout-letter|$M/themes/fallout-letter.png|600
e/outpost-spark|$M/themes/outpost-spark.png|300
e/fallout/logo|$M/themes/fallout/logo.png|700
e/fallout/hero|$M/themes/fallout/hero.webp|1676
e/fallout/banner|$M/themes/fallout/banner.png|1866
e/fallout/person|$M/themes/fallout/person.webp|640
e/fallout/comic|$M/themes/fallout/comic1.webp|640
e/htn/cloud|$M/themes/htn/cloud.webp|500
e/htn/mountain-left|$M/themes/htn/mountain-left.webp|1100
e/htn/mountain-front|$M/themes/htn/mountain-front.webp|700
e/htn/tree-row|$M/themes/htn/tree-row.webp|1400
e/htn/grass|$M/themes/htn/grass.webp|1736
e/htn/bush-light|$M/themes/htn/light-bush-1.webp|500
e/htn/bush-dark|$M/themes/htn/dark-bush-1.webp|500
e/htn/bg-left|$M/themes/htn/bg-left.webp|300
e/htn/bg-right|$M/themes/htn/bg-right.webp|300
e/htn/table|$M/themes/htn/picnic-table.webp|1500
e/htn/table-edge|$M/themes/htn/table-edge.webp|834
e/htn/sparkle|$M/themes/htn/sparkle1.webp|108
e/htn/cassette|$M/themes/htn/cassette-final.webp|500
e/htn/polaroid|$M/themes/htn/polaroid-final.webp|500
e/htn/vinyl|$M/themes/htn/vinyl-final.webp|500
# ---------- site screenshots ----------
s/fh27|$M/shot-fraserhacks.png|1400
s/fh26|$M/shot-fh26.png|1400
s/cs|$M/shot-cs.png|1400
s/dayoneof|$M/shot-dayoneof.png|1400
s/undercity|$M/shot-undercity.png|1400
s/fallout|$M/shot-fallout.png|1400
s/outpost|$M/shot-outpost.png|1400
s/horizons|$M/shot-horizons.png|1400
s/supercon|$M/shot-supercon.png|1400
s/submerge|$M/shot-submerge.jpg|1400
EOF

# sky photo and the plain assets
cp -n "$A/concept/assets/sky-blueskywithclouds.jpg" "$OUT/sky.jpg" 2>/dev/null || true
for f in hc-flag undercity-logo fallout-flag; do cp -n "$A/$M/$f.svg" "$OUT/e/$f.svg" 2>/dev/null || true; done
for f in outpost-bgcheck1 outpost-checkcorner outpost-stripe horizons-logo hackaday-logo undercity-projectsbanner; do cp -n "$A/$M/themes/$f.svg" "$OUT/e/$f.svg" 2>/dev/null || true; done
for f in top_left top_right bottom_left bottom_right; do cp -n "$A/$M/themes/fallout-border-$f.webp" "$OUT/e/fallout-$f.webp" 2>/dev/null || true; done

# Hack the North's game card, with its printed "1000+ hackers" painted over by a stretched blank strip of the label
g="$A/$M/themes/htn/gamecard-yellow.webp"
if [[ ! -f "$OUT/e/htn/gamecard.webp" || "$g" -nt "$OUT/e/htn/gamecard.webp" ]]; then
  convert "$g" \( "$g" -crop 30x290+64+168 +repage -resize '280x290!' \) -geometry +64+168 -composite -quality 85 "$OUT/e/htn/gamecard.webp"
  echo "built $OUT/e/htn/gamecard.webp"
fi

# their bush divider, without the soft shadow arcs around it
g="$A/$M/themes/htn/separator.webp"
if [[ ! -f "$OUT/e/htn/separator.webp" || "$g" -nt "$OUT/e/htn/separator.webp" ]]; then
  # the shadows are translucent and the bushes opaque, so keep only the near-opaque pixels
  convert "$g" \( +clone -alpha extract -level 80%,96% \) -compose copy_opacity -composite -trim +repage -resize 1200x -quality 82 "$OUT/e/htn/separator.webp"
  echo "built $OUT/e/htn/separator.webp"
fi

# fonts borrowed from the Fallout and Hack the North sites
mkdir -p static/fonts
cp -n "$A/$M/themes/fallout/Hells-Bells.otf" static/fonts/hells-bells.otf 2>/dev/null || true
cp -n "$A/$M/themes/htn/castledown-heavy.woff2" static/fonts/castledown-heavy.woff2 2>/dev/null || true
for w in Medium Bold; do cp -n "$A/$M/themes/htn/Satoshi-$w.woff2" "static/fonts/satoshi-${w,,}.woff2" 2>/dev/null || true; done
