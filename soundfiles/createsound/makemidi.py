from mido import MidiFile, MidiTrack, Message
import os

base = 0  # C-1
note_name = ""
sustain = 480
start_velocity = 64
end_velocity = 64

for i in range(127):
    mid = MidiFile()
    track = MidiTrack()
    mid.tracks.append(track)

    note = base + i
    track.append(Message("note_on", note=note, velocity=start_velocity, time=0))
    track.append(Message("note_off", note=note, velocity=end_velocity, time=sustain))

    if i % 12 == 0:
        note_name = "C"
    elif i % 12 == 1:
        note_name = "C#"
    elif i % 12 == 2:
        note_name = "D"
    elif i % 12 == 3:
        note_name = "D#"
    elif i % 12 == 4:
        note_name = "E"
    elif i % 12 == 5:
        note_name = "F"
    elif i % 12 == 6:
        note_name = "F#"
    elif i % 12 == 7:
        note_name = "G"
    elif i % 12 == 8:
        note_name = "G#"
    elif i % 12 == 9:
        note_name = "A"
    elif i % 12 == 10:
        note_name = "A#"
    elif i % 12 == 11:
        note_name = "B"
    else:
        note_name = "err"

    octave = (base + i) // 12 - 1
    if octave == 0:
        octave = -1
    else:
        octave -= 1

    os.makedirs("midinotes", exist_ok=True)
    mid.save(f"midinotes/{note_name}{octave}.mid")
