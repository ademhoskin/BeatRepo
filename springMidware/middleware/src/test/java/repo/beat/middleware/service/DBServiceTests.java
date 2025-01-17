package repo.beat.middleware.service;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.Date;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import repo.beat.middleware.model.BeatZip;

@SpringBootTest
public class DBServiceTests {
    
    @Autowired
    private ZipFileDBService zipFileDBService;

    @Test
    void testDBService() {
        // save
        String s3Link = "hello:make";
        String zipName = "hello";
        Long zipSize = 188888888883L;
        Date newDate = new Date();
        BeatZip beatZip = new BeatZip(zipName, zipSize, s3Link, newDate);
        BeatZip savedBeatZip = zipFileDBService.saveZipFile(beatZip);
        assertEquals(savedBeatZip.getZipName(), zipName);
        assertEquals(savedBeatZip.getZipSize(), zipSize);
        assertEquals(savedBeatZip.getS3Link(), s3Link);

        // get
        BeatZip retrievedBeatZip = zipFileDBService.getZipFileById(savedBeatZip.getId());
        assertEquals(retrievedBeatZip.getZipName(), zipName);
        assertEquals(retrievedBeatZip.getZipSize(), zipSize);
        assertEquals(retrievedBeatZip.getS3Link(), s3Link);


        // delete
        zipFileDBService.deleteZipFile(savedBeatZip.getId());
    }

}