mididir="midinotes"
outdir="../defaultpiano/"
sf2file="GM.sf2"

mkdir -p "$outdir"
mkdir -p "$mididir"

for f in "$mididir"/*.mid; do
    base="$(basename "${f%.mid}")"

    fluidsynth -ni -F "$outdir/$base.wav" -r 44100 "$sf2file" "$f"
    ffmpeg -y -i "$outdir/$base.wav" -codec:a libmp3lame -qscale:a 2 "$outdir/$base.mp3"
    rm "$outdir/$base.wav"
done
