"use client";

import { useState } from "react";
import { Scanner } from "@yudiel/react-qr-scanner";
import { useRouter, useSearchParams } from "next/navigation";

type AttendanceData = {
  Code: string;
  Name: string;
  NoTable: string;
  StatusCheckin: number;
};

type ApiResponse = {
  code: number;
  message: string;
  data: AttendanceData;
};

type Tab = "barcode" | "import";
type FacingMode = "user" | "environment";

type DetectedBarcode = {
  rawValue: string;
  format: string;
};

export default function ScannerClientPage() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const [activeTab, setActiveTab] = useState<Tab>("barcode");
  const [lastScan, setLastScan] = useState<AttendanceData | null>(null);
  const [manualCode, setManualCode] = useState("");
  const [cameraOn, setCameraOn] = useState(false);
  const [facingMode, setFacingMode] = useState<FacingMode>("environment");

  const eventName = searchParams.get("eventName") || "Event Check-in";
  const eventRecid = searchParams.get("event") || "";
const applyScanResult = async (codeValue: string) => {
  if (!codeValue) return;

  if (!eventRecid) {
    alert("Event belum dipilih.");
    return;
  }
  
  try {
    const res = await fetch("http://localhost:8080/api/app/attendance/scan", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        code: codeValue,
        event_recid: eventRecid,
      }),
    });

    console.log("status:", res.status);

    if (!res.ok) {
      const txt = await res.text();
      console.error("API error:", txt);
      alert("Gagal check-in: " + txt);
      return;
    }

    const json : ApiResponse = await res.json();
    console.log("json:", json);

    const d = json.data;
    setLastScan({
      Code: d.Code,
      Name: d.Name,
      NoTable: d.NoTable,
      StatusCheckin: d.StatusCheckin,
    });
  } catch (err) {
    console.error("fetch error:", err);
    alert("Tidak bisa konek ke server (Failed to fetch)");
  }
};

  const handleDetected = (codes: DetectedBarcode[]) => {
    if (!codes || codes.length === 0) return;
    const value = codes[0].rawValue;
    applyScanResult(value);
  };

  const handleScanError = (error: unknown) => {
    console.error(error);
  };

  const handleManualCheckin = async () => {
    if (!manualCode.trim()) return;

    await applyScanResult(manualCode.trim()); 

    setManualCode(""); 
  };

  const handleExcelUpload = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file) return;

    if (!eventRecid) {
        alert("Event belum dipilih.");
        return;
    }

    const formData = new FormData();
    formData.append("file", file);
    formData.append("event_recid", eventRecid);

    try {
        const res = await fetch("http://localhost:8080/api/app/attendance/import", {
        method: "POST",
        body: formData,
        });

        if (!res.ok) {
        const text = await res.text();
        console.error(text);
        alert("Gagal import attendance");
        return;
        }

        const json = await res.json();
        console.log("Import result:", json);
        alert("Import attendance sukses");
    } catch (err) {
        console.error(err);
        alert("Terjadi error saat upload file");
    } finally {
        e.target.value = "";
    }
  };


  const statusColor =
    lastScan?.StatusCheckin == 1
      ? "bg-red-500/15 text-red-300 border-red-500/40"
      : "bg-emerald-500/15 text-emerald-300 border-emerald-500/40";

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-b from-black via-zinc-900 to-amber-950 px-4 py-6">
      <main className="w-full max-w-sm rounded-3xl bg-black/80 border border-amber-500/30 shadow-xl shadow-amber-900/30 px-5 py-6 text-zinc-50 space-y-5">
        <header className="flex items-center justify-between gap-3">
          <div className="flex items-center gap-2">
            <div className="h-9 w-9 rounded-full bg-amber-500 flex items-center justify-center text-xs font-bold text-black shadow-lg">
              USD
            </div>
            <div className="text-xs leading-tight">
              <p className="font-semibold tracking-wide text-amber-300">Admin Check-in</p>
              <p className="text-[10px] text-zinc-300">Admin</p>
            </div>
          </div>
          <span className="rounded-full bg-zinc-900 px-3 py-1 text-[10px] text-zinc-300 border border-amber-500/40">
            Mode: Scanner
          </span>
        </header>

        <p className="text-center text-xs font-semibold text-amber-300">
          {eventName}
        </p>

        <div className="flex rounded-full bg-zinc-900/80 p-1 border border-amber-500/30 text-xs">
          <button
            onClick={() => setActiveTab("barcode")}
            className={`flex-1 rounded-full px-3 py-2 transition ${
              activeTab === "barcode"
                ? "bg-amber-500 text-black font-semibold shadow-md"
                : "text-zinc-300"
            }`}
          >
            Barcode
          </button>
          <button
            onClick={() => setActiveTab("import")}
            className={`flex-1 rounded-full px-3 py-2 transition ${
              activeTab === "import"
                ? "bg-amber-500 text-black font-semibold shadow-md"
                : "text-zinc-300"
            }`}
          >
            Import Excel
          </button>
        </div>

        {activeTab === "barcode" && (
          <section className="space-y-4">
            <div className="flex items-center justify-between text-[11px]">
              <div className="inline-flex items-center rounded-full bg-zinc-900/80 border border-amber-500/30 p-1">
                <button
                  onClick={() => setFacingMode("environment")}
                  className={`px-3 py-1 rounded-full text-xs ${
                    facingMode === "environment"
                      ? "bg-amber-500 text-black font-semibold"
                      : "text-zinc-300"
                  }`}
                >
                  Belakang
                </button>
                <button
                  onClick={() => setFacingMode("user")}
                  className={`px-3 py-1 rounded-full text-xs ${
                    facingMode === "user"
                      ? "bg-amber-500 text-black font-semibold"
                      : "text-zinc-300"
                  }`}
                >
                  Depan
                </button>
              </div>
              <button
                onClick={() => setCameraOn((prev) => !prev)}
                className={`px-3 py-1 rounded-full text-[11px] border ${
                  cameraOn
                    ? "bg-red-500/90 border-red-400 text-black"
                    : "bg-zinc-900 border-amber-500/40 text-zinc-200"
                }`}
              >
                {cameraOn ? "Matikan Kamera" : "Nyalakan Kamera"}
              </button>
            </div>

            <div className="relative h-80 w-full rounded-2xl overflow-hidden border border-amber-500/40 bg-black">
              {cameraOn && (
                <div className="absolute inset-0">
                  <Scanner
                    onScan={handleDetected}
                    onError={handleScanError}
                    constraints={{ facingMode }}
                    styles={{
                      container: { width: "100%", height: "100%" },
                      video: { width: "100%", height: "100%", objectFit: "cover" },
                    }}
                    scanDelay={400}
                  />
                </div>
              )}
              {!cameraOn && (
                <div className="absolute inset-0 flex flex-col items-center justify-center gap-2 text-center px-6">
                  <p className="text-[11px] uppercase tracking-[0.18em] text-amber-300">
                    Kamera Nonaktif
                  </p>
                  <p className="text-xs text-zinc-200">
                    Tekan tombol Nyalakan Kamera untuk mulai scan barcode.
                  </p>
                </div>
              )}
              <div className="absolute inset-6 border border-dashed border-amber-400/60 rounded-2xl pointer-events-none"></div>
              <div className="absolute bottom-2 w-full text-center">
                <p className="text-[11px] text-amber-300 tracking-widest">
                  {cameraOn ? "Arahkan barcode ke kotak" : "Kamera belum aktif"}
                </p>
              </div>
            </div>

            <div className="space-y-2 rounded-2xl bg-zinc-900/70 border border-amber-500/20 px-4 py-3">
              <p className="text-[11px] text-zinc-300 mb-1">Check-in Manual</p>
              <div className="flex items-center gap-2">
                <input
                  value={manualCode}
                  onChange={(e) => setManualCode(e.target.value)}
                  placeholder="Masukkan kode / no. undian"
                  className="flex-1 rounded-full bg-black/60 border border-zinc-700 px-3 py-2 text-xs text-zinc-100 placeholder:text-zinc-500 focus:outline-none focus:border-amber-400"
                />
                <button
                  onClick={handleManualCheckin}
                  className="rounded-full bg-amber-500 px-3 py-2 text-[11px] font-semibold text-black active:scale-95"
                >
                  Check-in
                </button>
              </div>
            </div>

            {lastScan && (
              <div className={`rounded-2xl border px-4 py-3 text-xs space-y-1 ${statusColor}`}>
                <div className="flex items-center justify-between">
                  <span className="text-[11px] uppercase tracking-[0.15em]">Hasil Terakhir</span>
                  <span className="text-[11px]">
                    Status:{" "}
                    <strong className="uppercase">
                      {lastScan.StatusCheckin == 1 ? "Sudah Check-in" : "Belum Check-in"}
                    </strong>
                  </span>
                </div>
                <div className="flex items-center justify-between">
                  <span>Nomor Undian</span>
                  <span className="font-mono text-[11px]">{lastScan.Code}</span>
                </div>
                <div className="flex items-center justify-between">
                  <span>Nama</span>
                  <span className="font-medium">{lastScan.Name}</span>
                </div>
                <div className="flex items-center justify-between">
                  <span>No. Meja</span>
                  <span className="font-semibold">{lastScan.NoTable}</span>
                </div>
              </div>
            )}
          </section>
        )}

        {activeTab === "import" && (
            <section className="space-y-4">
                {/* Informasi Import */}
                <div className="rounded-2xl bg-zinc-900/70 border border-amber-500/20 px-4 py-3 text-xs space-y-2">
                <p className="text-[11px] text-zinc-300">
                    Import daftar peserta / undian dari file Excel.
                </p>
                <p className="text-[11px] text-amber-400">
                    Gunakan template agar format data sesuai.
                </p>
                </div>

                {/* Tombol Download Template */}
                <button
                onClick={() => {
                    const link = document.createElement("a");
                    link.href = "/template/template_attendance.xlsx"; 
                    link.download = "template_attendance.xlsx";
                    link.click();
                }}
                className="w-full rounded-2xl bg-gradient-to-r from-amber-400 via-amber-500 to-orange-500 
                            px-4 py-3 text-xs font-semibold text-black shadow-md active:scale-[0.98] 
                            transition text-center"
                >
                 Download Template Excel
                </button>

                {/* Upload Excel */}
                <label className="w-full flex cursor-pointer items-center justify-between gap-3 rounded-2xl
                                border border-dashed border-amber-400/50 bg-zinc-950/60 px-4 py-3 text-xs text-amber-100">
                <div className="flex flex-col">
                    <span className="text-[11px] uppercase tracking-[0.15em] text-amber-400">
                    Upload Excel
                    </span>
                    <span className="text-xs text-zinc-200">Pilih file .xlsx / .xls</span>
                </div>

                <div className="flex h-10 w-10 items-center justify-center rounded-full bg-amber-500/90 
                                text-black text-lg shadow-md">
                    📂
                </div>

                <input
                    type="file"
                    accept=".xlsx,.xls"
                    className="hidden"
                    onChange={handleExcelUpload}
                />
                </label>
            </section>
        )}


        <div className="flex gap-2 mt-1">
          <button
            onClick={() => router.push("/")}
            className="flex-1 rounded-full border border-amber-500/60 bg-zinc-900 px-4 py-3 text-sm font-semibold text-zinc-100 shadow-md active:scale-[0.98] transition-transform"
          >
            Back to Homepage
          </button>
            <button
              onClick={() =>
                router.push(
                  `/doorprize/list/LIVE?event=${eventRecid}&eventName=${encodeURIComponent(eventName)}`
                )
              }
              className="flex-1 rounded-full bg-gradient-to-r from-amber-400 via-amber-500 to-orange-500 px-4 py-3 text-sm font-semibold text-black shadow-lg shadow-amber-900/40 active:scale-[0.98] transition-transform"
            >
              Doorprize
            </button>


        </div>
      </main>
    </div>
  );
}
