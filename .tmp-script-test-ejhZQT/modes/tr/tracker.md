# Mod: tracker — Başvuru Takipçisi

`data/applications.md` dosyasını oku ve görüntüle.

**Takip Formatı:**

```markdown
| # | Tarih | Şirket | Rol | Puan | Durum | PDF | Rapor | Notlar |
```

Olası durumlar: `Evaluated` / `Değerlendirildi` → `Applied` / `Başvuruldu` → `Responded` / `Yanıtlandı` → `Interview` / `Mülakat` → `Offer` / `Teklif` / `Rejected` / `Reddedildi` / `Discarded` / `Vazgeçildi` / `SKIP` / `ATLA`

- `Evaluated` / `Değerlendirildi` = teklif raporla değerlendirildi, karar bekleniyor
- `Applied` / `Başvuruldu` = aday başvurusunu gönderdi
- `Responded` / `Yanıtlandı` = Şirket yanıt verdi (henüz mülakat değil)
- `Interview` / `Mülakat` = aktif mülakat süreci
- `Offer` / `Teklif` = iş teklifi alındı
- `Rejected` / `Reddedildi` = şirket tarafından reddedildi
- `Discarded` / `Vazgeçildi` = aday tarafından reddedildi veya teklif kapandı
- `SKIP` / `ATLA` = uygun değil, başvurma

Kullanıcı bir durumu güncellemeyi isterse, ilgili satırı manuel olarak düzenlemek yerine `node set-status.mjs` betiğini kullanın.

Ayrıca istatistikleri de gösterin:
- Toplam başvuru sayısı
- Duruma göre dağılım
- Ortalama puan
- PDF oluşturulanların yüzdesi (%)
- Rapor oluşturulanların yüzdesi (%)
