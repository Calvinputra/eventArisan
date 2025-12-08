"use client";

import { useEffect, useRef, useState } from "react";
import { useRouter, useSearchParams } from "next/navigation";

type RawAttendanceFromApi = {
  recid: string;
  Code: string;
  Name: string;
  NoTable: number;
  StatusCheckin: number;
};

type Attendance = {
  AttendanceRecid: string;
  Code: string;
  Name: string;
  NoTable: number;
  StatusCheckin: number;
};

type AttendanceListResponse = {
  code: number;
  message: string;
  data: RawAttendanceFromApi[];
};

type DoorprizeApiItem = {
  eventRecid: string;
  attendanceRecid: string;
  attendanceName: string;
  attendanceCode: string;
};

type DoorprizeListResponse = {
  code: number;
  message: string;
  data: DoorprizeApiItem[];
};

type WinnerDisplay = {
  id: string;
  attendanceRecid: string;
  name: string;
  code: string;
  time?: string;
};

export default function DoorprizePage() {
  const router = useRouter();
  const searchParams = useSearchParams();

  const eventRecid = searchParams.get("event") || "";
  const eventName = searchParams.get("eventName") || "Doorprize Event";

  const [rotation, setRotation] = useState(0);
  const [isSpinning, setIsSpinning] = useState(false);

  const [allAttendees, setAllAttendees] = useState<Attendance[]>([]);
  const [remaining, setRemaining] = useState<Attendance[]>([]);
  const [winners, setWinners] = useState<WinnerDisplay[]>([]);
  const [currentWinner, setCurrentWinner] = useState<WinnerDisplay | null>(null);

  const [isLoading, setIsLoading] = useState(false);
  const [loadError, setLoadError] = useState<string | null>(null);

  const spinTimeoutRef = useRef<ReturnType<typeof setTimeout> | null>(null);

  const [attendeeMap, setAttendeeMap] = useState<Record<string, Attendance>>({});

  const fetchAttendees = async () => {
    if (!eventRecid) {
      setLoadError("Event belum dipilih.");
      return;
    }

    setIsLoading(true);
    setLoadError(null);

    try {
      const res = await fetch(
        `http://localhost:8080/api/app/attendance/list/LIVE?event=${eventRecid}`,
        { method: "GET" }
      );

      if (!res.ok) {
        setLoadError("Gagal mengambil data attendance.");
        return;
      }

      const json: AttendanceListResponse = await res.json();

      const mapped: Attendance[] = json.data.map((r) => ({
        AttendanceRecid: r.recid,
        Code: r.Code,
        Name: r.Name,
        NoTable: r.NoTable,
        StatusCheckin: r.StatusCheckin,
      }));

      const checkedIn = mapped;

      setAllAttendees(checkedIn);
      setRemaining(checkedIn);

      const map: Record<string, Attendance> = {};
      checkedIn.forEach((a) => {
        map[a.AttendanceRecid] = a;
      });
      setAttendeeMap(map);
    } catch (err) {
      setLoadError("Tidak bisa konek server attendance.");
    } finally {
      setIsLoading(false);
    }
  };

  const fetchDoorprizeWinners = async () => {
    try {
      const res = await fetch(
        `http://localhost:8080/api/app/doorprize/list/LIVE?event=${eventRecid}`,
        { method: "GET" }
      );

      if (!res.ok) {
        const txt = await res.text();
        console.error("Error fetch doorprize list:", txt);
        return;
      }

      const json: DoorprizeListResponse = await res.json();
      const winnerCodes = new Set(
        json.data.map((item) => item.attendanceRecid)
      );

      setRemaining((prev) =>
        prev.filter((a) => !winnerCodes.has(a.AttendanceRecid))
      );

      const list: WinnerDisplay[] = json.data.map((item, idx) => ({
        id: `${item.attendanceRecid}-${idx}`,
        code: item.attendanceCode,
        name: item.attendanceName,
        attendanceRecid: item.attendanceRecid,
      }));

      setWinners(list);
    } catch (err) {
      console.error("Error fetch doorprize winners:", err);
    }
  };

  const wheelItems = remaining.length > 0 ? remaining : allAttendees;

  useEffect(() => {
    fetchAttendees();
    return () => {
      if (spinTimeoutRef.current) clearTimeout(spinTimeoutRef.current);
    };
  }, [eventRecid]);

  useEffect(() => {
    if (Object.keys(attendeeMap).length > 0) {
      fetchDoorprizeWinners();
    }
  }, [attendeeMap]);

  const handleSpin = () => {
    if (isSpinning || remaining.length === 0) return;

    const segmentAngle = 360 / remaining.length;
    const winnerIndex = Math.floor(Math.random() * remaining.length);
    const winner = remaining[winnerIndex];

    const angleDeg = segmentAngle * winnerIndex;
    const desiredModulo = 360 - angleDeg;

    const base = rotation;
    let final = desiredModulo;
    while (final <= base + 1080) {
      final += 360;
    }

    setRotation(final);
    setIsSpinning(true);

    if (spinTimeoutRef.current) clearTimeout(spinTimeoutRef.current);

    spinTimeoutRef.current = setTimeout(() => {
      const time = new Date().toLocaleTimeString();

      const win: WinnerDisplay = {
        id: `${winner.Code}-${Date.now()}`,
        attendanceRecid: winner.AttendanceRecid,
        name: winner.Name,
        code: winner.Code,
        time,
      };

      setCurrentWinner(win);
      setRemaining((prev) =>
        prev.filter((x) => x.AttendanceRecid !== winner.AttendanceRecid)
      );
      setIsSpinning(false);

      (async () => {
        try {
          const res = await fetch(
            "http://localhost:8080/api/app/doorprize/create",
            {
              method: "POST",
              headers: {
                "Content-Type": "application/json",
              },
              body: JSON.stringify({
                eventRecid: eventRecid,
                attendanceRecid: winner.AttendanceRecid,
              }),
            }
          );

          if (!res.ok) {
            const txt = await res.text();
            console.error("Error create doorprize:", txt);
            alert("Gagal menyimpan pemenang ke server.");
            return;
          }

          await fetchDoorprizeWinners();
        } catch (err) {
          console.error("Error create doorprize:", err);
          alert("Tidak bisa konek ke server (doorprize/create).");
        }
      })();
    }, 3200);
  };

  const totalPeserta = allAttendees.length;
  const sudahMenang = winners.length;

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-b from-black via-zinc-900 to-[#3d1f0f] px-4 py-6">
      <main className="w-full max-w-sm rounded-3xl bg-black/70 border border-yellow-500/30 shadow-xl shadow-yellow-900/30 px-5 py-6 space-y-5">
        <header className="flex items-center justify-between text-yellow-300">
          <div>
            <p className="text-[10px] uppercase tracking-[0.25em]">Lucky Wheel</p>
            <p className="text-sm font-semibold">{eventName}</p>
          </div>
          <div className="text-right text-[11px] text-yellow-200">
            <p className="opacity-70">Peserta (checked-in)</p>
            <p className="font-semibold">
              {totalPeserta === 0
                ? "-"
                : `${totalPeserta - remaining.length}/${totalPeserta}`}
            </p>
            <p className="text-[10px] text-yellow-500">
              Pemenang: {sudahMenang} | Sisa code: {remaining.length}
            </p>
          </div>
        </header>

        {(isLoading || loadError) && (
          <div className="rounded-2xl bg-black/60 border border-yellow-500/40 px-4 py-3 text-[11px] text-yellow-200">
            {isLoading && <p>Memuat data peserta...</p>}
            {loadError && <p className="text-red-300">{loadError}</p>}
          </div>
        )}

        <section className="flex flex-col items-center gap-4">
          <div className="relative h-56 w-56 rounded-full bg-gradient-to-b from-yellow-500 via-yellow-600 to-amber-700 shadow-[0_20px_40px_rgba(255,200,0,0.15)] flex items-center justify-center border border-yellow-400/50">
            <div className="absolute -top-4 left-1/2 -translate-x-1/2 h-8 w-6 bg-yellow-100 rounded-b-full shadow-md flex items-center justify-center border border-yellow-400">
              <div className="h-3 w-3 bg-red-500 rounded-full shadow" />
            </div>

            <div
              className="relative h-48 w-48 rounded-full bg-black overflow-hidden transition-transform duration-[3200ms] ease-out border border-yellow-400/30"
              style={{ transform: `rotate(${rotation}deg)` }}
            >
              <div className="absolute inset-0 bg-gradient-to-b from-[#1f1a0a] to-[#3a260d]" />

              {wheelItems.map((att, i) => {
                const angle = (360 / wheelItems.length) * i;
                return (
                  <div
                    key={att.AttendanceRecid}
                    className="absolute inset-0 flex items-center justify-center"
                    style={{ transform: `rotate(${angle}deg)` }}
                  >
                    <div
                      className="flex items-center justify-center"
                      style={{ transform: "translateY(-64px) rotate(-90deg)" }}
                    >
                      <span className="text-[11px] font-extrabold text-yellow-200">
                        {att.Code}
                      </span>
                    </div>
                  </div>
                );
              })}

              <div className="absolute inset-10 rounded-full border-[6px] border-dotted border-yellow-400/70" />

              {currentWinner && (
                <div className="absolute inset-0 flex items-center justify-center pointer-events-none">
                  <div className="rounded-full bg-black/70 px-4 py-2 text-center border border-yellow-400/60">
                    <p className="text-[10px] text-yellow-300 uppercase tracking-[0.18em]">
                      Pemenang
                    </p>
                    <p className="text-xs font-bold text-yellow-100 mt-1">
                      {currentWinner.code}
                    </p>
                    <p className="text-[11px] text-yellow-200">
                      {currentWinner.name}
                    </p>
                  </div>
                </div>
              )}
            </div>

            <button
              onClick={handleSpin}
              disabled={isSpinning || remaining.length === 0 || !!loadError}
              className={`absolute h-20 w-20 rounded-full text-sm font-bold border-4 border-yellow-400 shadow-lg 
                ${
                  remaining.length === 0 || loadError
                    ? "bg-zinc-600 text-yellow-200"
                    : isSpinning
                    ? "bg-yellow-300 text-black"
                    : "bg-yellow-500 text-black hover:bg-yellow-400"
                }`}
            >
              {remaining.length === 0
                ? "Habis"
                : loadError
                ? "Error"
                : isSpinning
                ? "Spin..."
                : "GO"}
            </button>
          </div>

          <div className="w-full rounded-2xl bg-black/50 border border-yellow-500/30 shadow-inner px-4 py-3 text-xs text-yellow-200 space-y-1">
            <p className="text-[11px] font-semibold tracking-wide text-yellow-300">
              Winning Info
            </p>

            {currentWinner ? (
              <div className="space-y-1">
                <p className="text-sm font-bold text-yellow-100">
                  {currentWinner.name}
                </p>
                <p className="text-[11px] text-yellow-300">
                  No. Code: {currentWinner.code}
                </p>
                {currentWinner.time && (
                  <p className="text-[10px] text-yellow-500">
                    {currentWinner.time}
                  </p>
                )}
              </div>
            ) : (
              <p className="text-[11px] text-yellow-500">
                Tekan tombol GO untuk mengundi nomor code dari peserta yang sudah
                check-in.
              </p>
            )}
          </div>
        </section>

        <section className="space-y-3">
          <div className="flex items-center justify-between text-[11px] text-yellow-200">
            <p className="font-semibold flex items-center gap-1">
              <span>Daftar Pemenang</span>
              <span className="text-[10px] px-2 py-[2px] rounded-full bg-yellow-500/15 border border-yellow-500/40 text-yellow-200">
                {winners.length} orang
              </span>
            </p>
            {winners.length > 0 && (
              <p className="text-[10px] text-yellow-500">
                Terbaru di paling atas
              </p>
            )}
          </div>

          <div className="max-h-64 overflow-y-auto rounded-2xl bg-black/60 border border-yellow-500/30 text-yellow-200 text-xs shadow-inner">
            {winners.length === 0 && (
              <div className="px-4 py-4 text-[11px] text-yellow-500 text-center">
                Belum ada pemenang di table doorprize.
              </div>
            )}

            {winners.map((w, idx) => {
              const isTop = idx === 0;
              const isPodium = idx < 3;

              return (
                <div
                  key={w.id}
                  className={`px-4 py-3 flex items-center gap-3 border-b border-yellow-900/50 last:border-b-0
                    ${
                      isTop
                        ? "bg-gradient-to-r from-yellow-500/15 via-yellow-600/10 to-amber-600/10"
                        : "bg-black/40"
                    }`}
                >
                  <div
                    className={`h-7 w-7 rounded-full flex items-center justify-center text-[11px] font-bold
                      ${
                        isPodium
                          ? "bg-yellow-500 text-black"
                          : "bg-yellow-900/60 text-yellow-200"
                      }`}
                  >
                    {idx + 1}
                  </div>

                  <div className="flex-1 flex flex-col gap-1">
                    <div className="flex items-center justify-between gap-2">
                      <span className="text-sm font-semibold text-yellow-100 truncate max-w-[140px]">
                        {w.name}
                      </span>
                      <span className="text-[10px] px-2 py-[2px] rounded-full bg-yellow-500/15 border border-yellow-500/50 text-yellow-100">
                        {w.code}
                      </span>
                    </div>
                    {w.time && (
                      <div className="flex items-center justify-between text-[10px] text-yellow-500">
                        <span>Waktu: {w.time}</span>
                      </div>
                    )}
                  </div>
                </div>
              );
            })}
          </div>
        </section>

        <button
          onClick={() =>
            router.push(
              `/scanner?event=${eventRecid}&eventName=${encodeURIComponent(
                eventName
              )}`
            )
          }
          className="w-full mt-1 rounded-full bg-gradient-to-r from-amber-400 via-amber-500 to-orange-500 px-4 py-3 text-sm font-semibold text-black shadow-lg shadow-amber-900/40 active:scale-[0.98] transition-transform"
        >
          Scanner Page
        </button>
      </main>
    </div>
  );
}
