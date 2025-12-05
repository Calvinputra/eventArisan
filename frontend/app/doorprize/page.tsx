"use client";

import { useEffect, useRef, useState } from "react";
import { useRouter } from "next/navigation";
type Winner = {
  id: number;
  name: string;
  prize: string;
  time: string;
};

const prizes = [
  "Voucher 100K",
  "Voucher 50K",
  "Merchandise",
  "Mug Eksklusif",
  "Powerbank",
  "Payung",
  "Snack Box",
  "No Prize",
];

const sampleAttendees = [
  "Andi Wijaya",
  "Budi Santoso",
  "Citra Dewi",
  "Dewi Lestari",
  "Erik Gunawan",
  "Fajar Hidayat",
  "Grace Lim",
  "Hendra Kusuma",
];

export default function DoorprizePage() {
  const router = useRouter();
  const [rotation, setRotation] = useState(0);
  const [isSpinning, setIsSpinning] = useState(false);
  const [winners, setWinners] = useState<Winner[]>([]);
  const [remaining, setRemaining] = useState<string[]>(sampleAttendees);
  const [currentWinner, setCurrentWinner] = useState<Winner | null>(null);
  const spinTimeoutRef = useRef<ReturnType<typeof setTimeout> | null>(null);

  const handleSpin = () => {
    if (isSpinning || remaining.length === 0) return;

    const base = rotation % 360;
    const extra = 5 * 360;
    const random = Math.floor(Math.random() * 360);
    const final = base + extra + random;

    const randomIndex = Math.floor(Math.random() * remaining.length);
    const name = remaining[randomIndex];

    setRotation(final);
    setIsSpinning(true);

    if (spinTimeoutRef.current) clearTimeout(spinTimeoutRef.current);

    spinTimeoutRef.current = setTimeout(() => {
      const prize = prizes[Math.floor(Math.random() * prizes.length)];
      const time = new Date().toLocaleTimeString();

      const win: Winner = {
        id: Date.now(),
        name,
        prize,
        time,
      };

      setCurrentWinner(win);
      setWinners((prev) => [win, ...prev]);
      setRemaining((prev) => prev.filter((x) => x !== name));
      setIsSpinning(false);
    }, 3200);
  };

  useEffect(() => {
    return () => {
      if (spinTimeoutRef.current) clearTimeout(spinTimeoutRef.current);
    };
  }, []);

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-b from-black via-zinc-900 to-[#3d1f0f] px-4 py-6">
      <main className="w-full max-w-sm rounded-3xl bg-black/70 border border-yellow-500/30 shadow-xl shadow-yellow-900/30 px-5 py-6 space-y-5">
        
        {/* HEADER */}
        <header className="flex items-center justify-between text-yellow-300">
          <div>
            <p className="text-xs uppercase tracking-[0.25em]">Lucky Wheel</p>
            <p className="text-sm font-semibold">Doorprize Attendance</p>
          </div>
          <div className="text-right text-[11px] text-yellow-200">
            <p className="opacity-70">Checked-in</p>
            <p className="font-semibold">
              {sampleAttendees.length - remaining.length}/{sampleAttendees.length}
            </p>
          </div>
        </header>

        {/* WHEEL */}
        <section className="flex flex-col items-center gap-4">
          <div className="relative h-56 w-56 rounded-full bg-gradient-to-b from-yellow-500 via-yellow-600 to-amber-700 shadow-[0_20px_40px_rgba(255,200,0,0.15)] flex items-center justify-center border border-yellow-400/50">
            
            {/* POINTER */}
            <div className="absolute -top-4 left-1/2 -translate-x-1/2 h-8 w-6 bg-yellow-100 rounded-b-full shadow-md flex items-center justify-center border border-yellow-400">
              <div className="h-3 w-3 bg-red-500 rounded-full shadow" />
            </div>

            {/* WHEEL */}
            <div
              className="relative h-48 w-48 rounded-full bg-black overflow-hidden transition-transform duration-[3200ms] ease-out border border-yellow-400/30"
              style={{ transform: `rotate(${rotation}deg)` }}
            >
              <div className="absolute inset-0 bg-gradient-to-b from-[#1f1a0a] to-[#3a260d]" />

              {prizes.map((label, i) => {
                const angle = (360 / prizes.length) * i;
                return (
                  <div
                    key={label}
                    className="absolute inset-0 flex items-start justify-center"
                    style={{ transform: `rotate(${angle}deg)` }}
                  >
                    <div className="h-1/2 w-px bg-yellow-500/50" />
                    <div
                      className="absolute top-6 text-[10px] font-semibold text-yellow-200 text-center w-20"
                      style={{ transform: `rotate(90deg)` }}
                    >
                      {label}
                    </div>
                  </div>
                );
              })}

              <div className="absolute inset-6 rounded-full border-[6px] border-dotted border-yellow-400/70" />
            </div>

            {/* GO BUTTON */}
            <button
              onClick={handleSpin}
              disabled={isSpinning || remaining.length === 0}
              className={`absolute h-20 w-20 rounded-full text-sm font-bold border-4 border-yellow-400 shadow-lg 
                ${
                  remaining.length === 0
                    ? "bg-zinc-600 text-yellow-200"
                    : isSpinning
                    ? "bg-yellow-300 text-black"
                    : "bg-yellow-500 text-black hover:bg-yellow-400"
                }`}
            >
              {remaining.length === 0 ? "Habis" : isSpinning ? "Spin..." : "GO"}
            </button>
          </div>

          {/* WIN INFO */}
          <div className="w-full rounded-2xl bg-black/50 border border-yellow-500/30 shadow-inner px-4 py-3 text-xs text-yellow-200 space-y-1">
            <p className="text-[11px] font-semibold tracking-wide text-yellow-300">Winning Info</p>

            {currentWinner ? (
              <div className="space-y-1">
                <p className="text-sm font-bold text-yellow-100">{currentWinner.name}</p>
                <p className="text-[11px] text-yellow-300">Prize: {currentWinner.prize}</p>
                <p className="text-[10px] text-yellow-500">{currentWinner.time}</p>
              </div>
            ) : (
              <p className="text-[11px] text-yellow-500">
                Tekan tombol GO untuk mengundi pemenang dari peserta yang telah check-in.
              </p>
            )}
          </div>
        </section>

        {/* WINNERS LIST */}
        <section className="space-y-2">
          <div className="flex items-center justify-between text-[11px] text-yellow-200">
            <p className="font-semibold">Daftar Pemenang</p>
            <p className="opacity-80">{winners.length} pemenang</p>
          </div>

          <div className="max-h-64 overflow-y-auto rounded-2xl bg-black/50 border border-yellow-500/20 text-yellow-200 text-xs divide-y divide-yellow-900 shadow-inner">
            {winners.length === 0 && (
              <div className="px-4 py-3 text-[11px] text-yellow-500">
                Belum ada pemenang. Tekan tombol GO.
              </div>
            )}

            {winners.map((w) => (
              <div key={w.id} className="px-4 py-3 flex items-center justify-between gap-3">
                <div className="flex flex-col">
                  <span className="text-sm font-semibold text-yellow-100">{w.name}</span>
                  <span className="text-[11px] text-yellow-400">Prize: {w.prize}</span>
                </div>
                <span className="text-[10px] text-yellow-600">{w.time}</span>
              </div>
            ))}
          </div>
        </section>

        {/* BOTTOM BUTTON */}
        <button
          onClick={() => router.push("/scanner")}
          className="w-full mt-1 rounded-full bg-gradient-to-r from-amber-400 via-amber-500 to-orange-500 px-4 py-3 text-sm font-semibold text-black shadow-lg shadow-amber-900/40 active:scale-[0.98] transition-transform"
        >
          Scanner Page
        </button>

      </main>
    </div>
  );
}
