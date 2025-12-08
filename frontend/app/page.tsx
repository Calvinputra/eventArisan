"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";

type ApiEvent = {
  recid: string;
  Name: string;
  Description: string;
  Location: string;
  StartDateTime: number;
  Status: string;
};

type UiEvent = {
  id: string;
  name: string;
  date: string;
};

export default function Home() {
  const router = useRouter();
  const [events, setEvents] = useState<UiEvent[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchEvents = async () => {
      try {
        setLoading(true);
        setError(null);

        const res = await fetch("http://localhost:8080/api/app/event/list/LIVE");
        if (!res.ok) {
          throw new Error("Gagal ambil data event");
        }

        const json = await res.json();
        const data: ApiEvent[] = json.data || [];

        // Map bentuk API → bentuk yang dipakai UI
        const mapped: UiEvent[] = data.map((e, idx) => {
          let formattedDate = "-";
          if (e.StartDateTime && e.StartDateTime > 0) {
            formattedDate = new Date(e.StartDateTime * 1000).toLocaleDateString(
              "id-ID",
              { day: "numeric", month: "long", year: "numeric" }
            );
          }

          return {
            id: e.recid && e.recid !== "" ? e.recid : `event-${idx}`,
            name: e.Name || "(Tanpa nama)",
            date: formattedDate,
          };
        });

        setEvents(mapped);
      } catch (err) {
        console.error(err);
        setError("Tidak bisa memuat daftar event.");
      } finally {
        setLoading(false);
      }
    };

    fetchEvents();
  }, []);

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-b from-black via-zinc-900 to-[#3d1f0f] px-4 py-6">
      <main className="w-full max-w-sm rounded-3xl bg-black/70 border border-yellow-500/30 shadow-xl shadow-yellow-900/40 px-5 py-6 space-y-6 text-yellow-200">
        <h1 className="text-center text-lg font-semibold text-yellow-300 tracking-wide">
          Pilih Event
        </h1>

        <p className="text-center text-xs text-yellow-400">
          Silakan pilih event yang ingin dikelola check-in &amp; doorprize.
        </p>

        {loading && (
          <p className="text-center text-xs text-yellow-400">
            Memuat event...
          </p>
        )}

        {error && !loading && (
          <p className="text-center text-xs text-red-400">{error}</p>
        )}

        {!loading && !error && (
          <div className="space-y-3">
            {events.length === 0 && (
              <p className="text-center text-xs text-yellow-500">
                Belum ada event.
              </p>
            )}

            {events.map((ev) => (
              <button
                key={ev.id}
                onClick={() =>
                  router.push(
                    `/scanner?event=${encodeURIComponent(
                      ev.id
                    )}&eventName=${encodeURIComponent(ev.name)}`
                  )
                }
                className="w-full text-left px-4 py-3 rounded-2xl bg-black/60 border border-yellow-500/30 hover:bg-yellow-500/10 active:scale-[0.98] transition flex flex-col shadow-md"
              >
                <span className="text-sm font-semibold text-yellow-200">
                  {ev.name}
                </span>
                <span className="text-[11px] text-yellow-500">{ev.date}</span>
              </button>
            ))}
          </div>
        )}

        <p className="text-center text-[10px] text-yellow-500 pt-2">
          Admin Event Selection
        </p>
      </main>
    </div>
  );
}
