package repo.beat.middleware.service;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.http.ResponseEntity;

@SpringBootTest
public class S3ServiceTests {
    String bucketName = "cppbeatproj";
    String uploadObjectKey = "javatest.txt";
    String uploadFilePath =  "~/Workspace/BeatRepo/javatest.txt";
    String downloadFilePath = "~/Workspace/BeatRepo/javatest-reclaimed.txt";
    
    @Autowired
    private ZipFileS3Service zipFileS3Service;

    @Test
    void testS3Service() {
        ResponseEntity<String> resUp = zipFileS3Service.uploadZipFile(bucketName, uploadObjectKey, uploadFilePath);
        assertEquals(200, resUp.getStatusCode());

        // download
        ResponseEntity<String> resDown = zipFileS3Service.downloadZipFile(bucketName, uploadObjectKey, downloadFilePath);
        assertEquals(200, resDown.getStatusCode());

        // delete
        ResponseEntity<String> resDel = zipFileS3Service.deleteZipFile(bucketName, uploadObjectKey);
        assertEquals(200, resDel.getStatusCode());

    }
}
