package repo.beat.middleware.service;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import repo.beat.middleware.model.BeatZip;
import repo.beat.middleware.repository.ZipFileRepository;

@Service
public class ZipFileDBServiceImpl implements ZipFileDBService{
    @Autowired
    private ZipFileRepository zipFileRepository;

    @Override
    public BeatZip saveZipFile(BeatZip beatZip) {
        return zipFileRepository.save(beatZip);
    }

    @Override
    public List<BeatZip> getAllZipFiles() {
        return zipFileRepository.findAll();
    }

    @Override
    public BeatZip getZipFileById(Long id) {
        return zipFileRepository.findById(id).get();
    }

    @Override
    public void deleteZipFile(Long id) {
        zipFileRepository.deleteById(id);
    }
    
}
