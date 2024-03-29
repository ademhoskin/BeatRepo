package repo.beat.middleware.service;

import java.util.List;

import repo.beat.middleware.model.BeatZip;

interface ZipFileDBService  {
    BeatZip saveZipFile(BeatZip beatZip); // must upload as well?
    List<BeatZip> getAllZipFiles(); 
    BeatZip getZipFileById(Long id);
    void deleteZipFile(Long id);
}
